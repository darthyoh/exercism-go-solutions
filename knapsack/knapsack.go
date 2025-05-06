package knapsack

import "math"

type Item struct {
	Weight, Value int
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) int {
	if len(items) == 0 || maximumWeight == 0 {
		return 0
	}

	values := [][]float64{}

	//iterate over items
	for i, item := range items {
		valuesItem := []float64{}
		//iterate over possible weight
		for weight := 0; weight <= maximumWeight; weight++ {

			if item.Weight > weight { //item is too big for this weight
				if i > 0 { //indicate the precedent max value for this weight...
					valuesItem = append(valuesItem, values[i-1][weight])
				} else { //...or 0 if this is the first item
					valuesItem = append(valuesItem, 0)
				}

				continue
			}
			//item can be put so :
			if i == 0 { //indicate the item value for the first item...
				valuesItem = append(valuesItem, float64(item.Value))
			} else { //... or the max value between precedent value for this weight or item value + value of the precedent item for lesser weight
				valuesItem = append(valuesItem, math.Max(values[i-1][weight], float64(item.Value)+values[i-1][weight-item.Weight]))
			}

		}
		values = append(values, valuesItem)
	}

	return int(values[len(items)-1][maximumWeight])

}
