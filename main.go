package main

import (
	"flag"
	"fmt"
)

func main() {
	var fPath string

	// Get fpath to process by command line argument
	flag.StringVar(&fPath, "fPath", "", "File path to load the problem description from")
	flag.Parse()
	if len(fPath) == 0 {
		fmt.Errorf("you must specify a valid .json file path")
		return
	}

	ProblemToSolve, err := LoadProblemDefinitionFromJSON(fPath)
	if err != nil {
		fmt.Errorf("%s", err)
		return
	}
	maxValue, bestLoots := Knapsack(ProblemToSolve.Capacity, ProblemToSolve.Loots)
	fmt.Printf("Maximum value: %d, By selecting Loots: %s\n", maxValue, *bestLoots)
}
