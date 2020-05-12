package main

import "fmt"

type Treasure struct {
	weight int
	value  int
}

// Go through an array of Treasure, return the sum of the value.
// The array of Treasure is passed by reference to avoid passing heavy objects.
func getTotalValue(treasures *[]Treasure) int {
	var tot int = 0
	for _, t := range *treasures {
		tot += t.value
	}
	return tot
}

func Knapsack(capacity int, treasures []Treasure) (int, *[]Treasure) {
	var nTreasures int = len(treasures)
	var dp [][][]Treasure = make([][][]Treasure, nTreasures+1)
	dp[0] = make([][]Treasure, capacity+1)
	for i, t := range treasures {
		t_idx := i + 1
		dp[t_idx] = make([][]Treasure, capacity+1)
		row := make([][]Treasure, capacity+1)
		for w := 0; w < capacity+1; w++ {
			elem := make([]Treasure, nTreasures)
			if i == 0 || w == 0 {
				row[w] = elem
				continue
			}
			var maxValWithoutCurrentTreasure int = getTotalValue(&dp[t_idx-1][w])
			var maxValWithCurrentTreasure int = 0

			if w >= t.weight {
				var numTimesTFit int = w / t.weight
				maxValWithCurrentTreasure = t.weight * numTimesTFit
				var remainingCapa int = w - t.weight*numTimesTFit
				maxValWithCurrentTreasure += getTotalValue(&dp[t_idx-1][remainingCapa])

				if maxValWithCurrentTreasure > maxValWithoutCurrentTreasure {
					var addenda = make([]Treasure, numTimesTFit)
					for k := 0; k < numTimesTFit; k++ {
						addenda[k] = t
					}
					dp[t_idx][w] = append(dp[t_idx-1][remainingCapa], addenda...)
				} else {
					dp[t_idx][w] = dp[t_idx-1][w]
				}
			}
			row[w] = elem
		}
		dp[i] = row
	}
	var tot int = getTotalValue(&dp[nTreasures][capacity])

	return tot, &dp[nTreasures][capacity]
}

func main() {
	fmt.Printf("Hello Knapsack")
	var capa int = 5

	var t1 = Treasure{1, 1}
	var t2 = Treasure{2, 25}
	var t3 = Treasure{3, 30}
	var t5 = Treasure{5, 50}
	var treasures = []Treasure{t1, t2, t3, t5}

	const expected = 55

	var value_to_test, _ = Knapsack(capa, treasures)
	fmt.Printf("Result: %d, expected: %d", value_to_test, expected)
}
