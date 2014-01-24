package survey

import (
	"strings"
	"errors"
)

type QuestionType byte

const (
	RangeQuestionType QuestionType = iota
	ParaQuestionType QuestionType = iota
)

type Question struct {
	Title        string
	Section      int
	QuestionType QuestionType
}

func QuestionTypeFromString(s string) (QuestionType, error) {
	switch strings.ToLower(s) {
		case "1-5":
			return RangeQuestionType, nil
		case "para":
			return ParaQuestionType, nil
		default:
			return 0, errors.New("Invalid question type '" + s + "'")
	}
}

func QuestionTypeToString(b QuestionType) (string, error) {
	switch b {
		case RangeQuestionType:
			return "1-5", nil
		case ParaQuestionType:
			return "para", nil
		default:
			return "", errors.New("Invalid question type '" + string(b) + "'")
	}
}