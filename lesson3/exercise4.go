package main

import ()

func FiveSteps(array [5]int) [5]int {
	otv := [5]int{}
	for i := 0; i < len(array); i++ {
		otv[len(array) - i - 1] = array[i]
	}
	return otv
}