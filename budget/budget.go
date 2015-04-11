package budget

import (
	"time"
)

type money struct {
	notional float64
	currency string
}

const (
	USD string = "USD"
	HUF string = "HUF"
)


type statement struct {
	date time.Time
	cash, savings, debt money
}

func GetStatement(date time.Time) statement {
	s := statement{date, money{0,USD}, money{0,USD}, money{0,USD}}
	return s
}


