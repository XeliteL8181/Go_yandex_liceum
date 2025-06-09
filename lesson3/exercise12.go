package main

import (
	"errors"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if nums == nil {
		return nil, errors.New("входной слайс равен nil")
	}
	if n < 0 {
		return nil, errors.New("n не может быть меньше нуля")
	}

	result := []int{}

	for _, num := range nums {
		if num < limit {
			result = append(result, num)
			if len(result) == n {
				break
			}
		}
	}

	return result, nil
}
