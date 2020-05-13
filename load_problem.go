package knapsack

import (
	"encoding/json"
	"fmt"
)

type ProblemDefinition struct {
	Capacity int    `json:"capacity"`
	Loots    []Loot `json:"loots"`
}

func LoadProblemDefinitionFromJSON() ProblemDefinition {
	// Load a json file containing capacity and a list of Loot
	content := `{
  "capacity": 5,
  "loots": [
    {"weight":1, "value":1},
    {"weight":2, "value":25},
    {"weight":3, "value":30},
    {"weight":5, "value":50}
  ]
}`
	textBytes := []byte(content)

	problemDefinition1 := ProblemDefinition{}

	err := json.Unmarshal(textBytes, &problemDefinition1)

	if err != nil {
		fmt.Println(err)
		return ProblemDefinition{}
	}
	fmt.Println(problemDefinition1.Capacity)
	fmt.Println(problemDefinition1.Loots)
	return problemDefinition1
}
