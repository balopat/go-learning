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

func (s *MySuite) TestStatementDate(c *C) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	time := time.Date(2015, time.April, 10, 0, 0, 0, 0, loc)

	statement := GetStatement(time)

	c.Assert(statement.date, Equals, time)		

	c.Assert(statement.cash, Equals, 0.0)		
	c.Assert(statement.savings, Equals, 0.0)		
	c.Assert(statement.debt, Equals, 0.0)		
}

