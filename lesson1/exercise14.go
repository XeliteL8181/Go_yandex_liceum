package main

import (
	"fmt"
	"strings"
)

func main13(){
	var cnt, otv int
	var word, text string
	otv = 0
	fmt.Scanln(&cnt)
	fmt.Scanln(&word)

	for range cnt{
		fmt.Scanln(&text)
		if strings.ToLower(word) == strings.ToLower(text){
			otv++
		}
	}
	fmt.Println(otv)
}	