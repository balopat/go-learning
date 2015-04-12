package budget

import (
	"time"
	"fmt"
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
	date                time.Time
	cash, savings, debt money
}

type inventory struct {
	date                time.Time
	cash, savings, debt money
}

var inventories []inventory 

func ResetBudget() {
	inventories = []inventory{}
}

func GetStatement(date time.Time) statement {
	lastInventory := findInventoryInEffectFor(date)
	return lastInventory.toStatement()
}

func Inventory(inventory inventory) {
	inventories = append(inventories,inventory)
}


func (i inventory) toStatement() statement {
	return statement {
		i.date,
		i.cash,
		i.savings,
		i.debt}
}


func DateFor(year int, month time.Month, dayOfMonth int) time.Time {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	time := time.Date(year, month, dayOfMonth, 0, 0, 0, 0, loc)
	return time
}

func findInventoryInEffectFor(date time.Time) inventory {
	prevInventory := zeroInventory(date)

	for _, inv := range inventories {
		if (inv.date.After(date)) {
			return prevInventory
		} else {
			prevInventory = inv
		}
	}

	return prevInventory
}

func zeroInventory(date time.Time) inventory {
	return inventory{
		date:    date,
		cash:    money{0, USD},
		savings: money{0, USD},
		debt:    money{0, USD}}
}