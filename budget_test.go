package budget

import (
	"time"
	"testing"
)

func TestBudget(t *testing.T) {
	t := time.Date(2015, time.April, 10, 0, 0, 0, 0, LoadLocation("America/New_York"))

	statement := GetBudget.statement(t)

	if (statement.Date != t) {
		t.Error("Expected date on statement is ",t, " but got ", statement.Date)
	}
}