package main

import (
	"testing"
)

var loot1 = Loot{1, 1}
var loot2 = Loot{2, 25}
var loot3 = Loot{3, 30}
var loot4 = Loot{10, 60}
var loot5 = Loot{20, 100}
var loot6 = Loot{30, 120}
var loot7 = Loot{7, 6}
var loot8 = Loot{6, 7}

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

// Define test tables on which to iterate to test Knapsack
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

//def test_4():
//capa: int = 5
//
//t1 = Treasure(3, 1)
//t2 = Treasure(2, 2)
//treasures: List[Treasure] = [t1, t2]
//expected_treasures: List[Treasure] = [t2] * 2
//expected: int = 4
//value_to_test, set_to_test = max_possible_treasure_value(capa, treasures)
//assert value_to_test == expected
//assert set_to_test == expected_treasures
//
//
//def test_5():
//capa: int = 10
//
//t1 = Treasure(1, 2)
//t2 = Treasure(2, 2)
//t3 = Treasure(3, 3)
//treasures: List[Treasure] = [t1, t2, t3]
//expected_treasures: List[Treasure] = [t1] * 10
//expected: int = 20
//value_to_test, set_to_test = max_possible_treasure_value(capa, treasures)
//assert value_to_test == expected
//assert expected_treasures == set_to_test
//
//
//def test_6():
//# Empty set
//capa: int = 10
//
//t1 = Treasure(11, 2)
//
//treasures: List[Treasure] = [t1]
//expected_treasures: List[Treasure] = []
//expected: int = 0
//value_to_test, set_to_test = max_possible_treasure_value(capa, treasures)
//assert value_to_test == expected
//assert expected_treasures == set_to_test
//
//
//def test_7():
//# Big input
//capa: int = 100000
//
//t1 = Treasure(1, 4)
//t2 = Treasure(2, 2)
//t3 = Treasure(3, 4)
//
//treasures: List[Treasure] = [t1, t2, t3]
//expected_treasures: List[Treasure] = [t1] * capa
//expected: int = capa * t1.value
//value_to_test, set_to_test = max_possible_treasure_value(capa, treasures)
//assert value_to_test == expected
//assert expected_treasures == set_to_test
