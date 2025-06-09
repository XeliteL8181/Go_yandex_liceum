package main

import (
	"unicode"
)

func CheckOnlyASCII(s string) bool {
	for _, i := range s {
		if i > unicode.MaxASCII {
			return false
		}
	}
	return true
}