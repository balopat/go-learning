package budget

import (
	"time"
)

type statement struct {
	Date time.Time
}

func GetStatement(date time.Time) statement {
	s := statement{date}
	return s
}