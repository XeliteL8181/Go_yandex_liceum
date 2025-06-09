package main

import "fmt"

func main14() {
	var word, otv string
	fmt.Scanln(&word)
	for _, letter := range word {
		if letter!= 'a' && letter != 'о' {
			otv += string(letter)
		}
	}

	fmt.Println(otv)
}