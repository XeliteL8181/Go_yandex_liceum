package main

import(
	"fmt"
)

func main12(){
	var word string
	fmt.Scanln(&word)

	for (word != "да") && (word != "нет") && (word != "чёрный") && (word != "белый") {
		fmt.Println("Игра продолжается")
		fmt.Scanln(&word)
	}
	fmt.Println("Поражение")
}