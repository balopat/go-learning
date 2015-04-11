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

	assertStatementIsZeroAtDate(c, time, statement)
}



func (s *MySuite) TestBudgetInventory(c *C) {
	//inventoryDate := dateFor(2015, time.April, 10)
	afterInventoryDate := dateFor(2015, time.April, 12)
	//beforeInventoryDate := dateFor(2015, time.April, 9)	

	statement := GetStatement(afterInventoryDate)
	
	assertStatementIsZeroAtDate(c, afterInventoryDate, statement)
}

func assertStatementIsZeroAtDate(c *C, date time.Time, statement statement) {
	c.Assert(statement.date, Equals, date)		
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
