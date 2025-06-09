package main

import (
	"fmt"
	"unicode"
)

func checkPassword(password string) bool {
	if (hasMinimumLength(password, 8) && hasUpper(password)) {
		return true
	}
	return false
}

func hasMinimumLength(password string, minLen int) bool {
	if len(password) >= minLen {
		return true
	}
	return false
}

func hasUpper(password string) bool {
	for _, letter := range password{
		if unicode.IsUpper(letter){
			return true
		}
	}
	return false
}

func main5(){
	var password string
	fmt.Scanln(&password)
	fmt.Println(checkPassword(password))
}