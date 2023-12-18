package main

import "math"

// possiblePackCombinations defines all possible pack combinations in the range of smallest to largest pack size will help determine the optimum number of packs.
func possiblePackCombinations(packSizes []int) []int {
	var possiblePackCombinations []int

	smallest := math.MaxInt
	largest := math.MinInt
	for _, size := range packSizes {
		if size < smallest {
			smallest = size
		}
		if size > largest {
			largest = size
		}
	}

	for i := smallest; i <= largest; i += smallest {
		possiblePackCombinations = append(possiblePackCombinations, i)
	}

	return possiblePackCombinations
}
