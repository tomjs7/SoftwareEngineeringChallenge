package main

// handlePacks calculates the optimum number of packs based on the input provided.
func handlePacks(target int, sortedPackSizes []int, possiblePackSizes []int, combinations map[int][]int) []int {
	var packsRequired []int
	maxPackCount := 0

	if target == 0 {
		return []int{0}
	}

	for target > 0 {
		// If the number of items requested is greater than the largest pack size,
		// increase the maxPackCount by 1 and minus the largest pack size from the input.
		for target > possiblePackSizes[len(possiblePackSizes)-1] {
			target -= possiblePackSizes[len(possiblePackSizes)-1]
			maxPackCount++
		}
		var prevElement int
		for i, currentElement := range possiblePackSizes {
			if i == 0 {
				// First element: where the requested number of items is less than or equal to the current pack size.
				if target <= currentElement {
					packsRequired, target = handleOutput(combinations[currentElement], currentElement, packsRequired, target)
				}
			} else if i == len(possiblePackSizes)-1 {
				// Last element: where the requested number of items is within the range and greater than the previous element.
				if target <= currentElement && target > prevElement {
					packsRequired, target = handleOutput(combinations[currentElement], currentElement, packsRequired, target)
				}
			} else {
				// Every element inbetween the first and last: where the requested number of items is greater than the previous
				// element and less than or equal to the current element.
				prevElement = possiblePackSizes[i-1]
				if target > prevElement && target <= currentElement {
					packsRequired, target = handleOutput(combinations[currentElement], currentElement, packsRequired, target)
				}
			}
		}
	}

	// Add the number of maximum pack sizes based on the count.
	for i := 0; i < maxPackCount; {
		maxElement := sortedPackSizes[0]
		packsRequired = append(packsRequired, maxElement)
		i++
	}

	return packsRequired
}

// handleOutput processes the combination, updates the packsRequired and resets the target to 0 if the optimum number of packs required is achieved.
func handleOutput(combination []int, currentElement int, packsRequired []int, target int) ([]int, int) {
	value := combination
	packsRequired = append(packsRequired, value...)
	target = 0
	return packsRequired, target
}
