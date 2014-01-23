package survey

import (
	"time"
)

type Survey struct {
	Title string
	IsWriteOnce bool
	IsAnonymous bool
	ActiveAfter time.Time
	ActiveUntil time.Time
	Sections []*Section
	Questions []*Question
}
