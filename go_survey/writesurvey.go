package survey

import (
	"encoding/csv"
	"strconv"
)

// WriteSurvey writes survey sections and questions out to CSV
// [record-type, label, question-type, section-id]
func WriteSurvey(w *csv.Writer, sections []*Section, questions []*Question) error {
	for _, section := range sections {
		if err := writeSection(w, section); err != nil {
			return err
		}
	}

	for _, question := range questions {
		if err := writeQuestion(w, question); err != nil {
			return err
		}
	}

	w.Flush()

	return nil
}

func writeSection(w *csv.Writer, section *Section) error {
	return w.Write([]string {
		"section",
		section.Title,
		"",
		strconv.Itoa(section.Id),
	})
}

func writeQuestion(w *csv.Writer, question *Question) error {
	questionType, err := QuestionTypeToString(question.QuestionType)

	if err != nil {
		return err
	}

	return w.Write([]string {
		"question",
		question.Title,
		questionType,
		strconv.Itoa(question.Section),
	})
}