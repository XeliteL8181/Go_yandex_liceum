package main

import ()

func SwapKeysAndValues(m map[string]string) map[string]string {
	otv := map[string]string{}
	for key, value := range m {
		otv[value] = key
	}

	return otv
}