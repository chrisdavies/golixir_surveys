defmodule SurveyReaderTest do
  use ExUnit.Case, async: true

  test "records must be 4 columns" do
    assert_raise CSVError, fn ->
      SurveyReader.readcsv("section, hello")
    end
  end

  test "records must have numeric section ids" do
    assert_raise CSVError, fn ->
      SurveyReader.readcsv("section, hello, , one")
    end
  end

  test "questions must have a title" do 
    assert_raise CSVError, fn ->
      SurveyReader.readcsv("question, , para, 0")
    end
  end

  test "sections must have a title" do 
    assert_raise CSVError, fn ->
      SurveyReader.readcsv("section, , , 0")
    end
  end

  test "invalid question types raise errors" do
    assert_raise CSVError, fn ->
      SurveyReader.readcsv("question, hello, foo, 0")
    end
  end

  test "invalid record types raise errors" do
    assert_raise CSVError, fn ->
      SurveyReader.readcsv("foo, hello, 1-5, 0")
    end
  end

  test "section deserialzes properly" do
    {[section], _} = SurveyReader.readcsv("Section, bonjour, , 1")
    assert section == Section.new(title: "bonjour", id: 1)
  end

  test "question type is not case-sensitive" do
    {_, [question]} = SurveyReader.readcsv("Question, bonjour, Para, 1")
    assert question == Question.new(title: "bonjour", type: QuestionType.para, section: 1)
  end

  test "questions and sections deserialize properly" do
    record = """
section, hello, , 0
question, q-range, 1-5, 0
question, q-para, para, 0
"""
    {sections, questions} = SurveyReader.readcsv(record)
    
    expectedQuestions = [
      Question.new(title: "q-para", type: QuestionType.para, section: 0),
      Question.new(title: "q-range", type: QuestionType.range, section: 0)]
    
    compare_lists expectedQuestions, questions
    compare_lists [Section.new(title: "hello", id: 0)], sections
  end

  def compare_lists([expected | t1], [actual | t2]) do
    assert expected == actual
    compare_lists t1, t2
  end

  def compare_lists([], []) do 
  end
end
