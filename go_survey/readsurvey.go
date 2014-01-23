package survey

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"strings"
)

// ReadSurvey converts a CSV file of format
// [record-type, label, question-type, section-id]
// into a list of sections and questions
func ReadSurvey(r *csv.Reader) (sections []*Section, questions []*Question, err error) {
	r.TrimLeadingSpace = true

	for {
		line, err := r.Read()

		if err != nil && err != io.EOF {
			return nil, nil, err
		}

		if line == nil || len(line) == 0 {
			break
		}

		record, err := extractRecord(line)

		if err != nil {
			return nil, nil, err
		}

		switch rec := record.(type) {
		case *Section:
			sections = append(sections, rec)
		case *Question:
			questions = append(questions, rec)
		}
	}

	return sections, questions, nil
}

func extractRecord(line []string) (interface{}, error) {
	if len(line) != 4 {
		return csvError("Records must be in the format 'type, label, question-type, section-id'")
	}

	sectionId, err := strconv.Atoi(line[3])

	if err != nil {
		return csvError("Section ids must be numeric")
	}

	recordType := strings.ToLower(line[0])

	return newRecord(recordType, line[1], line[2], sectionId)
}

func newRecord(recordType string, label string, questionType string, sectionId int) (interface{}, error) {
	if recordType == "section" {
		return &Section{Title: label, Id: sectionId}, nil
	} else if recordType == "question" {
		questionType = strings.ToLower(questionType)

		switch questionType {
		case "1-5", "para":
			return &Question{Title: label, QuestionType: questionType, Section: sectionId}, nil
		default:
			return csvError("Invalid question type '" + questionType + "'")
		}
	}

	return csvError("Invalid record type '" + recordType + "'")
}

func csvError(s string) (interface{}, error) {
	return nil, errors.New(s)
}