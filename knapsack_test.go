package main

import (
	"testing"
)

var t1 = Loot{1, 1}
var t2 = Loot{2, 25}
var t3 = Loot{3, 30}
var l1 = Loot{10, 60}
var l2 = Loot{20, 100}
var l3 = Loot{30, 120}

// Helper function to compare slices of Loot.
func lootsEq(a, b []Loot) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	for idx, aItm := range a {
		if aItm != b[idx] {
			return false
		}
	}
	return true
}

func TestLootsEq(t *testing.T) {
	var a = []Loot{t1, t2}
	var b = []Loot{t1, t2}
	if lootsEq(a, b) == false {
		t.Error("Expected an equality of two equals loot slices.")
	}
}

func TestLootsIneq(t *testing.T) {
	var a = []Loot{t1, t2}
	var b = []Loot{t1, t3}
	if lootsEq(a, b) == true {
		t.Error("Expected an inequality of two equals loot slices.")
	}
}

// Define test tables on which to iterate to test Knapsack
var tables = []struct {
	loots         []Loot
	capacity      int
	expectedValue int
	expectedLoots []Loot
}{
	{
		[]Loot{t1, t2, t3},
		5,
		55,
		[]Loot{t2, t3},
	},
	{
		[]Loot{l1, l2, l3},
		50,
		300,
		[]Loot{l1, l1, l1, l1, l1},
	},
}

// Iterate over the test case defined in the test table
func TestKnapsack(t *testing.T) {
	for _, table := range tables {
		var valueToTest, setToTest = Knapsack(table.capacity, table.loots)
		if valueToTest != table.expectedValue {
			t.Errorf("Value Error - Expected %d, got %d", table.expectedValue, valueToTest)
		}
		if lootsEq(*setToTest, table.expectedLoots) == false {
			t.Errorf("Loots selection Error - Expected %s, got %s", table.expectedLoots, *setToTest)
		}
	}
}
