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

	recordType := strings.ToLower(line[0])
	title := strings.TrimSpace(line[1])
	questionType := line[2]
	sectionId, err := strconv.Atoi(line[3])

	if err != nil {
		return csvError("Section ids must be numeric")
	}

	if title == "" {
		return csvError("Title cannot be empty")
	}

	return newRecord(recordType, title, questionType, sectionId)
}

func newRecord(recordType string, title string, questionType string, sectionId int) (interface{}, error) {
	if recordType == "section" {
		return newSection(title, sectionId)
	} else if recordType == "question" {
		return newQuestion(title, questionType, sectionId)
	}

	return csvError("Invalid record type '" + recordType + "'")
}

func newSection(title string, sectionId int) (interface{}, error) {
	return &Section{Title: title, Id: sectionId}, nil
}

func newQuestion(title string, questionType string, sectionId int) (interface{}, error) {
	questionTypeByte, err := QuestionTypeFromString(questionType)

	if err != nil {
		return nil, err
	}

	return &Question{Title: title, QuestionType: questionTypeByte, Section: sectionId}, nil
}

func csvError(s string) (interface{}, error) {
	return nil, errors.New(s)
}