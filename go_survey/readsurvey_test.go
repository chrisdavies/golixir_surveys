package survey

import (
	"encoding/csv"
	"strings"
	"testing"
)

func Test_records_must_be_4_columns(t *testing.T) {
	_, _, err := ReadSurvey(rec("Section, Woirld"))
	if err == nil {
		t.Error("err was nil")
	}
}

func Test_records_must_have_numeric_section_id(t *testing.T) {
	_, _, err := ReadSurvey(rec("Section, Hello, , two"))
	if err == nil {
		t.Error("err was nil")
	}
}

func Test_title_is_not_blank(t *testing.T) {
	_, _, err := ReadSurvey(rec("section, , , 0"))
	if err == nil {
		t.Error("Section title should be required")
	}

	_, _, err = ReadSurvey(rec("question, , para, 0"))
	if err == nil {
		t.Error("Question title should be required")
	}
}

func Test_records_deserialize_sections(t *testing.T) {
	ss, _, err := ReadSurvey(rec("Section, Whazzup, , 0"))
	if recordIsNotEmpty(t, ss, err) {
		var section = ss[0]
		assert(section.Title == "Whazzup", "Section title was "+section.Title, t)
	}
}

func Test_valid_question_types_supported(t *testing.T) {
	assertQuestionType("Question, Whazzup, 1-5, 0", RangeQuestionType, t)
	assertQuestionType("Question, Whazzup, para, 0", ParaQuestionType, t)
}

func Test_invalid_record_types_cause_errors(t *testing.T) {
	_, _, err := ReadSurvey(rec("Foo, Hello, , 0"))
	if err == nil {
		t.Error("Foo should have been an unrecognized record type")
	}
}

func Test_invalid_question_types_cause_errors(t *testing.T) {
	_, _, err := ReadSurvey(rec("Question, Whazzup, foo, 0"))
	if err == nil {
		t.Error("Question type 'foo' should not be supported")
	}
}

func rec(recs string) *csv.Reader {
	return csv.NewReader(strings.NewReader(recs))
}

func assert(cond bool, str string, t *testing.T) {
	if !cond {
		t.Error(str)
	}
}

func assertQuestionType(record string, questionType QuestionType, t *testing.T) {
	_, qs, err := ReadSurvey(rec(record))
	if recordIsNotEmpty(t, qs, err) {
		var question = qs[0]
		assert(question.QuestionType == questionType, "Question type was " + string(question.QuestionType) + " but expected " + string(questionType), t)
	}
}

func recordIsNotEmpty(t *testing.T, arr interface{}, err error) bool {
	if err != nil {
		t.Error(err)
	} else if arr == nil {
		t.Error("result was nil")
	} else {
		return true
	}

	return false
}
