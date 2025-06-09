package main

import ()

func Join(nums1, nums2 []int) []int {
	otv := make([]int, len(nums1) + len(nums2), cap(nums1) + cap(nums2))
	copy(otv[:len(nums1)], nums1)
	copy(otv[len(nums1):], nums2)

	return  otv
}