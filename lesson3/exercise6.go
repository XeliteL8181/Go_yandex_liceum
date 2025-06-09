package main

import ()

func FindMinMaxInArray(array [10]int) (int, int) {
	maxArr := array[0]
	minArr := array[0]
	
	for i := 0; i < 10; i++ {
		if array[i] > maxArr {
			maxArr = array[i]
		} else if array[i] < minArr {
			minArr = array[i]
		}
	}

	return maxArr, minArr
}