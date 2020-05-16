package test

import (
	"github.com/fpaupier/knapsack/pkg/knapsack"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"testing"
)

var loot16 = knapsack.Loot{5, 50}

// Constant to access test data
const TestDataDir string = "data"

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
		expectedLoots    []knapsack.Loot
		expectedError    error
	}{
		{
			filepath.Join(testDataPath, "test1.json"),
			5,
			[]knapsack.Loot{loot1, loot2, loot3, loot16},
			nil,
		},
		{
			filepath.Join(testDataPath, "fakePath.json"),
			0,
			nil,
			&os.PathError{Op: "open", Path: filepath.Join(testDataPath, "fakePath.json"), Err: syscall.ENOENT},
		},
	}

	for _, testCase := range testTable {
		ProblemDefinitionToTest, err := knapsack.LoadProblemDefinitionFromJSON(testCase.fPath)
		if testCase.expectedError != nil {
			if err == nil {
				t.Errorf("Expected error %s got <nil>", testCase.expectedError)
			} else {
				if err.Error() != testCase.expectedError.Error() {
					t.Errorf("Expected error %s got %s", testCase.expectedError, err)
				} else {
					// We encountered the expected error, no need to further check the capacity and expected result
					continue
				}
			}
		}
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
