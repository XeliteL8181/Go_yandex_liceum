package lesson5

import "strings"

func DeleteVowels(s string) string {
	result := strings.Replace(s, "a", "", -1)
	result = strings.Replace(result, "e", "", -1)
	result = strings.Replace(result, "u", "", -1)
	result = strings.Replace(result, "i", "", -1)
	result = strings.Replace(result, "o", "", -1)
	result = strings.Replace(result, "A", "", -1)
	result = strings.Replace(result, "E", "", -1)
	result = strings.Replace(result, "U", "", -1)
	result = strings.Replace(result, "I", "", -1)
	result = strings.Replace(result, "O", "", -1)

	return result
}