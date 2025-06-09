package main

func FindMaxKey(m map[int]int) int {
	otv := -1
	for num, _ := range m {
		if num > otv {
			otv = num
		}
	}

	return otv
}
