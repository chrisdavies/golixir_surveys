package survey

import (
	"encoding/csv"
	"testing"
	"bytes"
)

func Test_sections_and_questions_are_written(t *testing.T) {
	out := bytes.Buffer{}
	
	sections := []*Section {
		&Section{Title: "hello", Id: 1},
		&Section{Title: "world", Id: 2},
	}

	questions := []*Question {
		&Question{Title: "qhello", QuestionType: "1-5", Section: 1},
		&Question{Title: "qworld", QuestionType: "para", Section: 2},
	}

	err := WriteSurvey(csv.NewWriter(&out), sections, questions)

	if err != nil {
		t.Error(err)
	}

	actual := out.String()
	expected := `section,hello,"",1
section,world,"",2
question,qhello,1-5,1
question,qworld,para,2
`

	if actual != expected {
		t.Error("Actual: >>" + actual + "<<")
	}
}
