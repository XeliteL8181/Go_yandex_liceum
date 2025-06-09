package main

import ()

func SliceCopy(nums []int) []int {
	otv := make([]int, len(nums), len(nums))
	return otv
}