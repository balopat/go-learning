package budget

import (
	"time"
	"testing"
	"fmt"
)

func TestBudget(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if (err!=nil) {
		fmt.Println(err)
	}

	time := time.Date(2015, time.April, 10, 0, 0, 0, 0, loc)
	
	statement := GetStatement(time)

	if (statement.Date != time) {
		t.Error("Expected date on statement is ",t, " but got ", statement.Date)
	}
}