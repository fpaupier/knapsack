package test

import (
	"github.com/fpaupier/knapsack/pkg/knapsack"
	"testing"
)

// illegal loots
var negativeValueLoot = knapsack.Loot{1, -1}
var negativeWeightLoot = knapsack.Loot{-1, 1}
var zeroWeightLoot = knapsack.Loot{0, 1}

// legal loots
var loot1 = knapsack.Loot{1, 1}
var loot2 = knapsack.Loot{2, 25}
var loot3 = knapsack.Loot{3, 30}
var loot4 = knapsack.Loot{10, 60}
var loot5 = knapsack.Loot{20, 100}
var loot6 = knapsack.Loot{30, 120}
var loot7 = knapsack.Loot{7, 6}
var loot8 = knapsack.Loot{6, 7}
var loot9 = knapsack.Loot{3, 1}
var loot10 = knapsack.Loot{2, 2}
var loot11 = knapsack.Loot{3, 3}
var loot12 = knapsack.Loot{1, 2}
var loot13 = knapsack.Loot{11, 2}
var loot14 = knapsack.Loot{1, 4}
var loot15 = knapsack.Loot{3, 4}

// Helper function to compare slices of Loot.
func lootsEq(a, b []knapsack.Loot) bool {

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
	var a = []knapsack.Loot{loot1, loot2}
	var b = []knapsack.Loot{loot1, loot2}
	if lootsEq(a, b) == false {
		t.Error("Expected an equality of two equals loot slices.")
	}
}

func TestLootsInequality(t *testing.T) {
	var a = []knapsack.Loot{loot1, loot2}
	var b = []knapsack.Loot{loot1, loot3}
	if lootsEq(a, b) == true {
		t.Error("Expected an inequality of two equals loot slices.")
	}
}

// Test Knapsack implementation on various use cases
func TestKnapsack(t *testing.T) {
	// Define test tables on which to iterate to test Knapsack
	var bigCapacity = 1000
	var expectedLootsBigCapacityTestCase []knapsack.Loot
	for j := 0; j < bigCapacity; j++ {
		expectedLootsBigCapacityTestCase = append(expectedLootsBigCapacityTestCase, loot14)
	}
	var tables = []struct {
		loots         []knapsack.Loot
		capacity      int
		expectedValue int
		expectedLoots []knapsack.Loot
	}{
		{
			[]knapsack.Loot{loot1, loot2, loot3},
			5,
			55,
			[]knapsack.Loot{loot2, loot3},
		},
		{
			[]knapsack.Loot{negativeWeightLoot},
			5,
			0, // best thing to do here is to not take the loot at all
			nil,
		},
		{
			[]knapsack.Loot{negativeValueLoot},
			5,
			0, // best thing to do here is to not take the loot at all
			nil,
		},
		{
			[]knapsack.Loot{negativeValueLoot, loot1, loot2, loot3},
			5,
			55,
			[]knapsack.Loot{loot2, loot3},
		},
		{
			[]knapsack.Loot{zeroWeightLoot, loot1, loot2, loot3},
			5,
			55,
			[]knapsack.Loot{loot2, loot3},
		},
		{
			[]knapsack.Loot{negativeWeightLoot, loot1, loot2, loot3},
			5,
			55,
			[]knapsack.Loot{loot2, loot3},
		},
		{
			[]knapsack.Loot{loot1, loot2, loot3},
			-10, // Knapsack with negative capacity should return immediately a 0 maxValue
			0,
			nil,
		},
		{
			[]knapsack.Loot{loot1, loot2},
			0, // Knapsack with 0 capacity should return immediately a 0 maxValue
			0,
			nil,
		},
		{
			[]knapsack.Loot{},
			50,
			0,
			nil,
		},
		{
			[]knapsack.Loot{loot4, loot5, loot6},
			50,
			300,
			[]knapsack.Loot{loot4, loot4, loot4, loot4, loot4},
		},
		{
			[]knapsack.Loot{loot7, loot8},
			10,
			7,
			[]knapsack.Loot{loot8},
		},
		{
			[]knapsack.Loot{loot9, loot10},
			5,
			4,
			[]knapsack.Loot{loot10, loot10},
		},
		{
			[]knapsack.Loot{loot10, loot11, loot12},
			10,
			20,
			[]knapsack.Loot{loot12, loot12, loot12, loot12, loot12, loot12, loot12, loot12, loot12, loot12},
		},
		{
			[]knapsack.Loot{loot13},
			10,
			0,
			nil,
		},
		{
			[]knapsack.Loot{loot14, loot10, loot15},
			bigCapacity,
			bigCapacity * loot14.Value,
			expectedLootsBigCapacityTestCase,
		},
	}

	// Iterate over the test case defined in the test table
	for _, table := range tables {
		var valueToTest, setToTest = knapsack.Knapsack(table.capacity, table.loots)
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
