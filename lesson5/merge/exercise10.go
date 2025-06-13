package main

import (
	"slices"
)	

func SortAndMerge(left, right []int) []int {
	var result []int
	i, j := 0, 0

	slices.Sort(left)
	slices.Sort(right)

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}