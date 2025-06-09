package main

import ()

func CountingSort(contacts []string) map[string]int {
	otv := map[string]int{}
	for i := 0; i < len(contacts); i++ {
		name := contacts[i]
		_, found := otv[name]
		if !found {
			otv[name] = 1
		} else {
			otv[name] += 1
		}
	}

	return otv
}