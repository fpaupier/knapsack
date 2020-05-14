package knapsack

import (
	"encoding/json"
	"fmt"
)

type ProblemDefinition struct {
	Capacity int    `json:"capacity"`
	Loots    []Loot `json:"loots"`
}

// Load a json file containing a capacity and a list of Loot.
func LoadProblemDefinitionFromJSON(content string) (ProblemDefinition, error) {
	textBytes := []byte(content)
	problemDefinitionRecovered := ProblemDefinition{}
	err := json.Unmarshal(textBytes, &problemDefinitionRecovered)
	if err != nil {
		fmt.Println(err)
		return ProblemDefinition{}, err
	}
	return problemDefinitionRecovered, nil
}
