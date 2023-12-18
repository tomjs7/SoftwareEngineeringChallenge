package main

import "sort"

// sortPackSizes sorts the pack sizes from lowest to highest value helps to ensure consistency in the future if the packSizes array changes.
func sortPackSizes(packSizes []int) []int {
	sort.Slice(packSizes, func(i int, j int) bool {
		return packSizes[i] < packSizes[j]
	})

	return packSizes
}
