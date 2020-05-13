package knapsack

import (
	"testing"
)

// illegal loots
var negativeValueLoot = Loot{1, -1}

// legal loots
var loot1 = Loot{1, 1}
var loot2 = Loot{2, 25}
var loot3 = Loot{3, 30}
var loot4 = Loot{10, 60}
var loot5 = Loot{20, 100}
var loot6 = Loot{30, 120}
var loot7 = Loot{7, 6}
var loot8 = Loot{6, 7}
var loot9 = Loot{3, 1}
var loot10 = Loot{2, 2}
var loot11 = Loot{3, 3}
var loot12 = Loot{1, 2}
var loot13 = Loot{11, 2}
var loot14 = Loot{1, 4}
var loot15 = Loot{3, 4}

// Helper function to compare slices of Loot.
func lootsEq(a, b []Loot) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}
	lenA := len(a)
	if lenA != len(b) {
		return false
	}

	// Edge case of empty list that are equal.
	if lenA == 0 {
		return true
	}

	for idx, aItm := range a {
		if aItm != b[idx] {
			return false
		}
	}
	return true
}

func TestLootsEq(t *testing.T) {
	var a = []Loot{loot1, loot2}
	var b = []Loot{loot1, loot2}
	if lootsEq(a, b) == false {
		t.Error("Expected an equality of two equals loot slices.")
	}
}

func TestLootsIneq(t *testing.T) {
	var a = []Loot{loot1, loot2}
	var b = []Loot{loot1, loot3}
	if lootsEq(a, b) == true {
		t.Error("Expected an inequality of two equals loot slices.")
	}
}

// Test Knapsack implementation on various use cases
func TestKnapsack(t *testing.T) {
	// Define test tables on which to iterate to test Knapsack
	var bigCapacity = 1000
	var expectedLootsBigCapacityTestCase []Loot
	for j := 0; j < bigCapacity; j++ {
		expectedLootsBigCapacityTestCase = append(expectedLootsBigCapacityTestCase, loot14)
	}
	var tables = []struct {
		loots         []Loot
		capacity      int
		expectedValue int
		expectedLoots []Loot
	}{
		{
			[]Loot{loot1, loot2, loot3},
			5,
			55,
			[]Loot{loot2, loot3},
		},
		{
			[]Loot{negativeValueLoot, loot1, loot2, loot3},
			5,
			55,
			[]Loot{loot2, loot3},
		},
		{
			[]Loot{loot1, loot2, loot3},
			-10, // Knapsack with negative capacity should return immediately a 0 maxValue
			0,
			nil,
		},
		{
			[]Loot{loot1, loot2},
			0, // Knapsack with 0 capacity should return immediately a 0 maxValue
			0,
			nil,
		},
		{
			[]Loot{},
			50,
			0,
			nil,
		},
		{
			[]Loot{loot4, loot5, loot6},
			50,
			300,
			[]Loot{loot4, loot4, loot4, loot4, loot4},
		},
		{
			[]Loot{loot7, loot8},
			10,
			7,
			[]Loot{loot8},
		},
		{
			[]Loot{loot9, loot10},
			5,
			4,
			[]Loot{loot10, loot10},
		},
		{
			[]Loot{loot10, loot11, loot12},
			10,
			20,
			[]Loot{loot12, loot12, loot12, loot12, loot12, loot12, loot12, loot12, loot12, loot12},
		},
		{
			[]Loot{loot13},
			10,
			0,
			nil,
		},
		{
			[]Loot{loot14, loot10, loot15},
			bigCapacity,
			bigCapacity * loot14.value,
			expectedLootsBigCapacityTestCase,
		},
	}

	// Iterate over the test case defined in the test table
	for _, table := range tables {
		var valueToTest, setToTest = Knapsack(table.capacity, table.loots)
		if valueToTest != table.expectedValue {
			t.Errorf("Value Error - Expected %d, got %d", table.expectedValue, valueToTest)
		}
		if setToTest == nil {
			if table.expectedLoots != nil {
				t.Errorf("Loots selection Error - Expected %s, got <nil>", table.expectedLoots)
			} else {
				continue
			}
		}
		if lootsEq(*setToTest, table.expectedLoots) == false {
			t.Errorf("Loots selection Error - Expected %s, got %s", table.expectedLoots, *setToTest)
		}
	}
}
