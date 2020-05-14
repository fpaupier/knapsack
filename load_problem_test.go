package knapsack

import "testing"

var loot16 = Loot{5, 50}

func TestLoadProblemDefinitionFromJSON(t *testing.T) {

	// Define test table on which to iterate
	var testTable = []struct {
		rawContent       string
		expectedCapacity int
		expectedLoots    []Loot
		expectedError    error
	}{
		{
			`{  "capacity": 5,
						  "loots": [
							{"weight":1, "value":1},
							{"weight":2, "value":25},
							{"weight":3, "value":30},
							{"weight":5, "value":50}
						  ]
						}`,
			5,
			[]Loot{loot1, loot2, loot3, loot16},
			nil,
		},
	}

	for _, testCase := range testTable {
		ProblemDefinitionToTest, err := LoadProblemDefinitionFromJSON(testCase.rawContent)
		if err != testCase.expectedError {
			t.Errorf("Expected error %s got %s", testCase.expectedError, err)
		}
		if err == nil {
			if ProblemDefinitionToTest.Capacity != testCase.expectedCapacity {
				t.Errorf("Expected capacity of %d got %d", testCase.expectedCapacity, ProblemDefinitionToTest.Capacity)
			}

			if lootsEq(ProblemDefinitionToTest.Loots, testCase.expectedLoots) == false {
				t.Errorf("Expected loots: %s, got %s", testCase.expectedLoots, ProblemDefinitionToTest.Loots)
			}
		}
	}
}
