package survey

import (
	"errors"
	"strconv"
)

func ValidateSurvey(s *Survey) error {
	if s.ActiveAfter.IsZero() {
		return required("Active after")
	}

	if s.ActiveUntil.IsZero() {
		return required("Active until")
	}

	if s.Title == "" {
		return required("Title")
	}

	if s.Questions == nil {
		return required("Questions")
	}

	if s.Sections == nil {
		return required("Sections")
	}

	if s.ActiveAfter.After(s.ActiveUntil) {
		return errors.New("Active after must be an earlier date than active until")
	}

	return validate_questions(s)
}

func required(name string) error {
	return errors.New(name + " is required")
}

func validate_questions(s *Survey) error {
	var m = make(map[int]int)

	for _, sect := range s.Sections {
		m[sect.Id] = sect.Id
	}

	for _, q := range s.Questions {
		if _, ok := m[q.Section]; !ok {
			return errors.New("Question '" + q.Title + "' references unknown section '" + strconv.Itoa(q.Section) + "'")
		}
	}

	return nil
}