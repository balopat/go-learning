package budget

import (
	//"fmt"
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

func (s *MySuite) SetUpTest(c *C) {
    ResetBudget()
}

var _ = Suite(&MySuite{})

func (s *MySuite) TestEmptyBudget(c *C) {	
	time := DateFor(2015, time.April, 10)

	c.Assert(GetStatement(time), Equals,
		statement{
			date:    time,
			cash:    money{0, USD},
			savings: money{0, USD},
			debt:    money{0, USD}})
}

func (s *MySuite) TestBudgetInventory(c *C) {
	inventoryDate := DateFor(2015, time.April, 10)
	afterInventoryDate := DateFor(2015, time.April, 12)
	beforeInventoryDate := DateFor(2015, time.April, 9)

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
			date:    afterInventoryDate,
			cash:    money{3.4, USD},
			savings: money{1424224, USD},
			debt:    money{100, USD}})
}

func (s *MySuite) TestBudgetInventorySorting(c *C) {
	inventoryDate := DateFor(2015, time.April, 8)
	secondInventoryDate := DateFor(2015, time.April, 10)
	betweenInventoriesDate := DateFor(2015, time.April, 9)
	afterSecondInventoryDate := DateFor(2015, time.April,11)

	Inventory(inventory{
		date:    inventoryDate,
		cash:    money{3.4, USD},
		savings: money{1424224, USD},
		debt:    money{100, USD}})


	Inventory(inventory{
		date:    secondInventoryDate,
		cash:    money{1, USD},
		savings: money{2, USD},
		debt:    money{3, USD}})


	c.Assert(GetStatement(betweenInventoriesDate), Equals,
		statement{
			date:    betweenInventoriesDate,
			cash:    money{3.4, USD},
			savings: money{1424224, USD},
			debt:    money{100, USD}})

	c.Assert(GetStatement(afterSecondInventoryDate), Equals,
		statement{
			date:    afterSecondInventoryDate,
			cash:    money{1, USD},
			savings: money{2, USD},
			debt:    money{3, USD}})
}


