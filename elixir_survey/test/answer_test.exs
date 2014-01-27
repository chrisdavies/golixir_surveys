defmodule AnswerTest do
  use ExUnit.Case, async: true

  test "answers must not be fewer than questions" do
    should_err []
  end

  test "answers must not be greater than questions" do
    should_err [3, "stuff", "more stuff"]
  end

  test "valid answers should not err" do
    validate_answers [3, "stuff"]
  end

  test "range should be greater than -1" do
    should_err [-1, "stuff"]
  end

  test "range should be less than 6" do
    should_err [6, "stuff"]
  end

  test "range can be nil" do
    validate_answers [nil, "stuff"]
  end

  test "para can be nil" do
    validate_answers [2, nil]
  end

  test "range must be numeric" do
    should_err ["hello", "world"]
  end

  test "para must be string" do 
    should_err [3, 4]
  end

  def validate_answers(answers) do
      questions = [
        Question.new(title: "qhello", type: :range, section: 1),
        Question.new(title: "qworld", type: :para, section: 2)]

      AnswerValidator.validate(questions, answers)
  end

  def should_err(answers) do
    assert_raise ValidationError, fn ->
      validate_answers answers
    end
  end

end
