package main 

import ()

func DeleteLongKeys(m map[string]int) map[string]int {
	otv := map[string]int{}
	for key, value := range m {
		if len(key) >= 6 {
			otv[key] = value
		} 
	}

	return otv
}