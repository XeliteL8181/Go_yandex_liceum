package main

import ()

func SumOfValuesInMap(m map[int]int) int {
	otv := 0
	for _, num := range m {
		otv += num
	}

	return otv
}