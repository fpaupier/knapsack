package knapsack

import (
	"path/filepath"
	"runtime"
	"testing"
)

var loot16 = Loot{5, 50}

// Constant to access test data
const TestDataDir string = "test_data"

// Util function to get current working directory.
func getCurDir() string {
	//ex, err := os.Executable()
	_, curFilename, _, _ := runtime.Caller(1)
	curDir := filepath.Dir(curFilename)
	return curDir
}

var curDir = getCurDir()
var testDataPath = filepath.Join(curDir, TestDataDir)

func TestLoadProblemDefinitionFromJSON(t *testing.T) {

	// Define test table on which to iterate
	var testTable = []struct {
		fPath            string
		expectedCapacity int
		expectedLoots    []Loot
		expectedError    error
	}{
		{
			filepath.Join(testDataPath, "test1.json"),
			5,
			[]Loot{loot1, loot2, loot3, loot16},
			nil,
		},
	}

	for _, testCase := range testTable {
		ProblemDefinitionToTest, err := LoadProblemDefinitionFromJSON(testCase.fPath)
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
