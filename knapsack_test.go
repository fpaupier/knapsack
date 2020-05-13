package main

import (
	"sort"
	"testing"
)

var t1 = Loot{1, 1}
var t2 = Loot{2, 25}
var t3 = Loot{3, 30}
var t5 = Loot{5, 50}

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

func TestKnapsackExpectedValueNominal(t *testing.T) {
	var capa = 5

	var treasures = []Loot{t1, t2, t3, t5}

	const expectedValue = 55

	var valueToTest, _ = Knapsack(capa, treasures)

	if valueToTest != expectedValue {
		t.Errorf("Expected %d, got %d", expectedValue, valueToTest)
	}
}

func TestKnapsackExpectedSetNominal(t *testing.T) {
	var capa = 5

	var treasures = []Loot{t1, t2, t3, t5}

	var _, setToTest = Knapsack(capa, treasures)
	var expectedSet = []Loot{t2, t3}
	sort.Sort(ByWeight(*setToTest))

	if lootsEq(expectedSet, *setToTest) == false {
		t.Errorf("Expected %s, got %s", expectedSet, setToTest)
	}
}

func TestKnapsackFullOneItem(t *testing.T) {
	var capacity = 50

	var l1 = Loot{10, 60}
	var l2 = Loot{20, 100}
	var l3 = Loot{30, 120}

	var treasures = []Loot{l1, l2, l3}

	var valueToTest, setToTest = Knapsack(capacity, treasures)
	const expectedValue = 300

	var expectedSet = []Loot{l1, l1, l1, l1, l1}
	sort.Sort(ByWeight(*setToTest))

	if lootsEq(expectedSet, *setToTest) == false {
		t.Errorf("Expected set %s, got %s", expectedSet, *setToTest)
	}
	if expectedValue != valueToTest {
		t.Errorf("Expected value %d, got %d", expectedValue, valueToTest)
	}
}
