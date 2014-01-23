package survey

import (
	"testing"
	"time"
)

type ModSurvey func(*Survey)

func Test_valid_surveys_validate(t *testing.T) {
	shouldnt_err(t, func(s *Survey) { })
}

func Test_active_after_is_required(t *testing.T) {
	should_err(t, func(s *Survey) { s.ActiveAfter = time.Time{} });
}

func Test_active_until_is_required(t *testing.T) {
	should_err(t, func(s *Survey) { s.ActiveUntil = time.Time{} });
}

func Test_title_is_required(t *testing.T) {
	should_err(t, func(s *Survey) { s.Title = "" });
}

func Test_questions_are_required(t *testing.T) {
	should_err(t, func(s *Survey) { s.Questions = nil });
}

func Test_sections_are_required(t *testing.T) {
	should_err(t, func(s *Survey) { s.Sections = nil });
}

func Test_active_after_cant_be_greater_than_active_until(t *testing.T) {
	should_err(t, func(s *Survey) { s.ActiveUntil = s.ActiveUntil.AddDate(-1, 0, 0) });
}

func Test_questions_can_reference_existing_sections(t *testing.T) {
	shouldnt_err(t, func (s *Survey) {
		q := &Question{Title: "Yo", QuestionType: "1-5", Section: 1}
		sect := &Section{Title: "Bo", Id: 1}
		s.Questions = []*Question{ q }
		s.Sections = []*Section { sect }
	})
}

func Test_questions_cannot_reference_bad_sections(t *testing.T) {
	should_err(t, func (s *Survey) {
		q := &Question{Title: "Yo", QuestionType: "1-5", Section: 1}
		sect := &Section{Title: "Bo", Id: 0}
		s.Questions = []*Question{ q }
		s.Sections = []*Section { sect }
	})
}

func shouldnt_err(t *testing.T, fn ModSurvey) {
	survey := valid_survey()
	fn(survey)
	err := ValidateSurvey(survey)
	if err != nil {
		t.Error(err)
	}
}

func should_err(t *testing.T, fn ModSurvey) {
	survey := valid_survey()
	fn(survey)
	err := ValidateSurvey(survey)
	if err == nil {
		t.Error("err was nil")
	}
}

func valid_survey() *Survey {
	return &Survey{
		Title: "hello", 
		ActiveAfter: time.Now(), 
		ActiveUntil: time.Now(), 
		Sections: []*Section{}, 
		Questions: []*Question{}}
}