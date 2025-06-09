package main

import ()

func SumOfArray(array [6]int) int {
	otv := 0
	for i := 0; i < len(array); i++ {
		otv += array[i]
	}

	return otv
}