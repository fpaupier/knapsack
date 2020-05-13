package knapsack

import (
	"fmt"
)

// Loot represents an object we can add to the bag. A loot has a Weight and a Value.
// Weight must be strictly positive; Weight > 0.
// Value must be positive; Value >= 0.
// In a knapsack problem, we want to select the Loot so that they maximize the total Value.
type Loot struct {
	Weight int `json:"weight"`
	Value  int `json:"value"`
}

// Implements the fmt.String interface to get the representation of a Loot as a string.
func (t Loot) String() string {
	return fmt.Sprintf("(Weight: %d, Value: %d)", t.Weight, t.Value)
}

// ByWeight implements sort.Interface for []Loot based on the Weight field.
type ByWeight []Loot

func (a ByWeight) Len() int           { return len(a) }
func (a ByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByWeight) Less(i, j int) bool { return a[i].Weight < a[j].Weight }

// Go through an array of Loot, return the sum of the Value.
// The array of Loot is passed by reference to avoid passing heavy objects.
func getTotalValue(treasures *[]Loot) int {
	var tot = 0
	for _, t := range *treasures {
		tot += t.Value
	}
	return tot
}

// Compute the best Value possible to store within a Knapsack of a given `capacity`.
func Knapsack(capacity int, loots []Loot) (int, *[]Loot) {
	// Edge case of negative capacity, immediately return
	if capacity <= 0 {
		return 0, nil
	}
	var nLoots = len(loots)

	// Edge case of empty loots, immediately return
	if nLoots == 0 {
		return 0, nil
	}

	// Solve Knapsack using dynamic programming, dp is our 2D DP array
	var dp = make([][][]Loot, nLoots+1)
	dp[0] = make([][]Loot, capacity+1)
	for i, t := range loots {
		tIdx := i + 1
		dp[tIdx] = make([][]Loot, capacity+1)
		for w := 0; w < capacity+1; w++ {
			if w == 0 {
				continue
			}
			var maxValWithoutCurrentTreasure = getTotalValue(&dp[tIdx-1][w])
			var maxValWithCurrentTreasure = 0

			if w >= t.Weight {
				// Handle edge case of negative weights
				if t.Weight <= 0 {
					dp[tIdx][w] = dp[tIdx-1][w]
					continue
				}
				var numTimesTFit = w / t.Weight
				maxValWithCurrentTreasure = t.Value * numTimesTFit
				var remainingCapacity = w - t.Weight*numTimesTFit
				maxValWithCurrentTreasure += getTotalValue(&dp[tIdx-1][remainingCapacity])

				if maxValWithCurrentTreasure > maxValWithoutCurrentTreasure {
					var addenda = make([]Loot, numTimesTFit)
					for k := 0; k < numTimesTFit; k++ {
						addenda[k] = t
					}
					dp[tIdx][w] = append(dp[tIdx-1][remainingCapacity], addenda...)
				} else {
					dp[tIdx][w] = dp[tIdx-1][w]
				}
			}
		}
	}
	var tot = getTotalValue(&dp[nLoots][capacity])

	return tot, &dp[nLoots][capacity]
}
