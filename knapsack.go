package knapsack

import (
	"fmt"
	"sort"
	"time"
)

// Loot represents an object we can add to the bag. A loot has a weight and a value.
// In a knapsack problem, we want to select the Loot so that they maximize the total value.
type Loot struct {
	weight int
	value  int
}

// Implements the fmt.String interface to get the representation of a Loot as a string.
func (t Loot) String() string {
	return fmt.Sprintf("(weight: %d, value: %d)", t.weight, t.value)
}

// ByWeight implements sort.Interface for []Loot based on the weight field.
type ByWeight []Loot

func (a ByWeight) Len() int           { return len(a) }
func (a ByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByWeight) Less(i, j int) bool { return a[i].weight < a[j].weight }

// Go through an array of Loot, return the sum of the value.
// The array of Loot is passed by reference to avoid passing heavy objects.
func getTotalValue(treasures *[]Loot) int {
	var tot = 0
	for _, t := range *treasures {
		tot += t.value
	}
	return tot
}

// Compute the best value possible to store within a Knapsack of a given `capacity`.
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
		t_idx := i + 1
		dp[t_idx] = make([][]Loot, capacity+1)
		for w := 0; w < capacity+1; w++ {
			if w == 0 {
				continue
			}
			var maxValWithoutCurrentTreasure = getTotalValue(&dp[t_idx-1][w])
			var maxValWithCurrentTreasure = 0

			if w >= t.weight {
				var numTimesTFit = w / t.weight
				maxValWithCurrentTreasure = t.value * numTimesTFit
				var remainingCapacity = w - t.weight*numTimesTFit
				maxValWithCurrentTreasure += getTotalValue(&dp[t_idx-1][remainingCapacity])

				if maxValWithCurrentTreasure > maxValWithoutCurrentTreasure {
					var addenda = make([]Loot, numTimesTFit)
					for k := 0; k < numTimesTFit; k++ {
						addenda[k] = t
					}
					dp[t_idx][w] = append(dp[t_idx-1][remainingCapacity], addenda...)
				} else {
					dp[t_idx][w] = dp[t_idx-1][w]
				}
			}
		}
	}
	var tot = getTotalValue(&dp[nLoots][capacity])

	return tot, &dp[nLoots][capacity]
}

func main() {
	fmt.Printf("Hello Knapsack\n")
	var capa = 5

	var t1 = Loot{1, 1}
	var t2 = Loot{2, 25}
	var t3 = Loot{3, 30}
	var t5 = Loot{5, 50}
	var treasures = []Loot{t1, t2, t3, t5}

	const expectedValue = 55
	var expectedSet = []Loot{t2, t3}
	start := time.Now()
	var valueToTest, treasuresToTest = Knapsack(capa, treasures)
	end := time.Now()
	sort.Sort(ByWeight(*treasuresToTest))
	fmt.Printf("Knapsack computation took %s\n", end.Sub(start))
	fmt.Printf("Value: %d\nExpected Value: %d\n", valueToTest, expectedValue)
	fmt.Printf("Set: %s\nExpected Set: %s\n", expectedSet, expectedSet)

}
