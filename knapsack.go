package main

import "fmt"

type Treasure struct {
	weight int
	value  int
}

// Go through an array of Treasure, return the sum of the value.
// The array of Treasure is passed by reference to avoid passing heavy objects.
func getTotalValue(treasures *[]Treasure) int {
	var tot = 0
	for _, t := range *treasures {
		tot += t.value
	}
	return tot
}

func Knapsack(capacity int, treasures []Treasure) (int, *[]Treasure) {
	var nTreasures = len(treasures)
	var dp = make([][][]Treasure, nTreasures+1)
	dp[0] = make([][]Treasure, capacity+1)
	for i, t := range treasures {
		t_idx := i + 1
		dp[t_idx] = make([][]Treasure, capacity+1)
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
					var addenda = make([]Treasure, numTimesTFit)
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
	var tot = getTotalValue(&dp[nTreasures][capacity])

	return tot, &dp[nTreasures][capacity]
}

func main() {
	fmt.Printf("Hello Knapsack")
	var capa = 5

	var t1 = Treasure{1, 1}
	var t2 = Treasure{2, 25}
	var t3 = Treasure{3, 30}
	var t5 = Treasure{5, 50}
	var treasures = []Treasure{t1, t2, t3, t5}

	const expected = 55

	var valueToTest, _ = Knapsack(capa, treasures)
	fmt.Printf("Result: %d, expected: %d", valueToTest, expected)
}
