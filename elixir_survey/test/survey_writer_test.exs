defmodule SurveyWriterTest do
  use ExUnit.Case, async: true

  test "records are converted to csv" do
    actual = SurveyWriter.writecsv({
      [
        Section.new(title: "hello", id: 1),
        Section.new(title: "world", id: 2)],
      [
        Question.new(title: "qhello", type: :range, section: 1),
        Question.new(title: "qworld", type: :para, section: 2)]
    })

    expected = """
section,hello,,1
section,world,,2
question,qhello,1-5,1
question,qworld,para,2
"""
    assert expected == actual <> "\n"
  end
end
