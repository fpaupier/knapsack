package knapsack

import "testing"

func TestLoadProblemDefinitionFromJSON(t *testing.T) {
	var content string = `{
  "capacity": 5,
  "loots": [
    {"weight":1, "value":1},
    {"weight":2, "value":25},
    {"weight":3, "value":30},
    {"weight":5, "value":50}
  ]
}`
	ProblemDefinition1 := LoadProblemDefinitionFromJSON(content)
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
