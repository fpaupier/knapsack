package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ProblemDefinition struct {
	Capacity int    `json:"capacity"`
	Loots    []Loot `json:"loots"`
}

// Load a json file containing a capacity and a list of Loot.
func LoadProblemDefinitionFromJSON(fPath string) (ProblemDefinition, error) {
	textBytes, err := ioutil.ReadFile(fPath)
	if err != nil {
		return ProblemDefinition{}, err
	}
	problemDefinitionRecovered := ProblemDefinition{}
	err = json.Unmarshal(textBytes, &problemDefinitionRecovered)
	if err != nil {
		fmt.Println(err)
		return ProblemDefinition{}, err
	}
	return problemDefinitionRecovered, nil
}
