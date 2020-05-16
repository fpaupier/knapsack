package main

import (
	"flag"
	"fmt"
	"github.com/fpaupier/knapsack/pkg/knapsack"
)

func main() {
	var fPath string

	// Get fPath to process by command line argument
	flag.StringVar(&fPath, "fPath", "", "File path to load the problem description from")
	flag.Parse()
	if len(fPath) == 0 {
		fmt.Println("you must specify a valid .json file path")
		return
	}

	ProblemToSolve, err := knapsack.LoadProblemDefinitionFromJSON(fPath)
	if err != nil {
		fmt.Printf("got error: %s", err)
		return
	}
	maxValue, bestLoots := knapsack.Knapsack(ProblemToSolve.Capacity, ProblemToSolve.Loots)
	fmt.Printf("Maximum value: %d, By selecting Loots: %s\n", maxValue, *bestLoots)
}
