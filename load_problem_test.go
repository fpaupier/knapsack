package knapsack

import "testing"

func TestLoadProblemDefinitionFromJSON(t *testing.T) {
	ProblemDefinition1 := LoadProblemDefinitionFromJSON()
	if ProblemDefinition1.Capacity != 5 {
		t.Errorf("Expected capacity of 5 got %d", ProblemDefinition1.Capacity)
	}
	expectedLoots := []Loot{
		{
			Weight: 1,
			Value:  1,
		},
		{
			Weight: 2,
			Value:  25,
		},
		{
			Weight: 3,
			Value:  30,
		},
		{
			Weight: 5,
			Value:  50,
		},
	}
	if lootsEq(ProblemDefinition1.Loots, expectedLoots) == false {
		t.Errorf("Expected loots: %s, got %s", expectedLoots, ProblemDefinition1.Loots)
	}
}
