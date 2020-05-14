package knapsack

import (
	"encoding/json"
	"fmt"
)

type ProblemDefinition struct {
	Capacity int    `json:"capacity"`
	Loots    []Loot `json:"loots"`
}

func LoadProblemDefinitionFromJSON(content string) ProblemDefinition {
	// Load a json file containing capacity and a list of Loot
	textBytes := []byte(content)

	problemDefinitionRecovered := ProblemDefinition{}

	err := json.Unmarshal(textBytes, &problemDefinitionRecovered)

	if err != nil {
		fmt.Println(err)
		return ProblemDefinition{}
	}
	return problemDefinitionRecovered
}
