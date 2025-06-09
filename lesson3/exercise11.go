package main

import ()

func Mix(nums []int) []int {
	n := len(nums) / 2
	otv := make([]int, 0, 2 * n)
	for i := 0; i < n; i++ {
		otv = append(otv, nums[i])
		otv = append(otv, nums[n + i])
	}

	return otv
}