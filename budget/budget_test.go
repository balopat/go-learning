package budget

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestEmptyBudget(c *C) {
	time := dateFor(2015, time.April, 10)

	c.Assert(GetStatement(time), Equals,
		statement{
			date:    time,
			cash:    money{0, USD},
			savings: money{0, USD},
			debt:    money{0, USD}})
}

func (s *MySuite) TestBudgetInventory(c *C) {
	inventoryDate := dateFor(2015, time.April, 10)
	afterInventoryDate := dateFor(2015, time.April, 12)
	beforeInventoryDate := dateFor(2015, time.April, 9)

	Inventory(inventory{
		date:    inventoryDate,
		cash:    money{3.4, USD},
		savings: money{1424224, USD},
		debt:    money{100, USD}})

	c.Assert(GetStatement(beforeInventoryDate), Equals,
		statement{
			date:    beforeInventoryDate,
			cash:    money{0, USD},
			savings: money{0, USD},
			debt:    money{0, USD}})

	c.Assert(GetStatement(inventoryDate), Equals,
		statement{
			date:    inventoryDate,
			cash:    money{3.4, USD},
			savings: money{1424224, USD},
			debt:    money{100, USD}})

	c.Assert(GetStatement(afterInventoryDate), Equals,
		statement{
			date:    inventoryDate,
			cash:    money{3.4, USD},
			savings: money{1424224, USD},
			debt:    money{100, USD}})
}

func dateFor(year int, month time.Month, dayOfMonth int) time.Time {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	time := time.Date(year, month, dayOfMonth, 0, 0, 0, 0, loc)
	return time
}
