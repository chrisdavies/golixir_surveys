defmodule SurveyTest do
  use ExUnit.Case, async: true

  test "valid surveys do not throw exceptions" do
    validate_survey (fn(survey) -> survey end)
  end

  test "active after is required" do
    should_err fn(survey) -> survey.active_after(nil) end
  end

  test "active until is required" do
    should_err fn(survey) -> survey.active_until(nil) end
  end

  test "questions cannot be nil" do
    should_err fn(survey) -> survey.questions(nil) end
  end

  test "sections cannot be nil" do
    should_err fn(survey) -> survey.sections(nil) end
  end

  test "title is required" do
    should_err fn(survey) -> survey.title(nil) end
  end

  test "active after can't be greater than active until" do
    past = Calendar.plus(Calendar.universal_time(), days: -2)
    should_err fn(survey) -> survey.active_until(past) end
  end

  test "questions can reference an existing section" do
    validate_survey fn(survey) ->
      ss = [Section.new(title: "hello", id: 1)]
      qs = [Question.new(title: "bonjour", type: "para", section: 1)]
      survey.questions(qs).sections(ss)
    end
  end

  test "questions can't reference a non-existent section" do
    should_err fn(survey) ->
      ss = [Section.new(title: "hello", id: 0)]
      qs = [Question.new(title: "bonjour", type: "para", section: 1)]
      survey.questions(qs).sections(ss)
    end
  end

  def validate_survey(func) do
      now = Calendar.universal_time()
      survey = Survey.new(title: "hello", active_after: now, active_until: Calendar.plus(now, days: 1))
      SurveyValidator.validate(func.(survey))
  end

  def should_err(func) do
    assert_raise ValidationError, fn ->
      validate_survey func
    end
  end

end
