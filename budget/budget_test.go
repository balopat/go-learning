package budget

import (
	"fmt"
	"testing"
	"time"
    . "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestEmptyBudget(c *C) {
	time := dateFor(2015, time.April, 10)

	statement := GetStatement(time)

	c.Assert(statement.date, Equals, time)		

	c.Assert(statement.cash, Equals, money{0,USD})		
	c.Assert(statement.savings, Equals, money{0,USD})		
	c.Assert(statement.debt, Equals, money{0,USD})		
}

func dateFor(year int, month time.Month, dayOfMonth int) time.Time {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	time := time.Date(year, month, dayOfMonth, 0, 0, 0, 0, loc)
	return time
}
