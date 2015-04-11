package budget

import (
	"time"
)

type statement struct {
	date time.Time
	cash, savings, debt float64
}

func GetStatement(date time.Time) statement {
	s := statement{date, 0, 0, 0}
	return s
}
