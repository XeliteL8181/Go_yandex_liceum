package main

import (
	"fmt"
	"unicode"
)

func ratePassword(password string) string{
	rate := 0
	if hasMinimumLength1(password, 8){
		rate++
	}
	if hasUpper1(password){
		rate++
	}
	if hasLowerCase(password){
		rate++
	}

	switch{
	case rate == 1:
		return "слабый пароль"
	case rate == 2:
		return "средний пароль"
	case rate == 3:
		return "сложный пароль"
	default:
		return "придумайте новый пароль, данный пароль не обладает достаточной безопасностью"
	}
}

func hasMinimumLength1(password string, minLen int) bool {
	if len(password) >= minLen {
		return true
	}
	return false
}

func hasUpper1(password string) bool {
	for _, letter := range password{
		if unicode.IsUpper(letter){
			return true
		}
	}
	return false
}

func hasLowerCase(password string) bool {
	for _, letter := range password{
		if unicode.IsLower(letter){
			return true
		}
	}
	return false
}

func main6(){
	var password string
	fmt.Scanln(&password)
	fmt.Println(ratePassword(password))
}