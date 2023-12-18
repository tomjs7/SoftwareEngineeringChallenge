package main

import "sort"

// findCombinations takes the pack sizes and possible combinations as inputs and outputs a map where:
// Keys of the map are integers representing possible pack combinations
// Values represents the optimum number of packs required to make the pack combination
// Using a map here allows to quickly lookup the combination of packs needed for a given pack size.
func findCombinations(packSizes []int, possiblePackSizes []int) map[int][]int {
	combinations := make(map[int][]int)

	for _, num := range possiblePackSizes {
		combinationsArr := make([]int, 0)
		remaining := num

		sort.Slice(packSizes, func(i int, j int) bool {
			return packSizes[i] > packSizes[j]
		})

		for _, value := range packSizes {
			for remaining >= value {
				combinationsArr = append(combinationsArr, value)
				remaining -= value
			}
		}
		combinations[num] = combinationsArr
	}

	return combinations
}
