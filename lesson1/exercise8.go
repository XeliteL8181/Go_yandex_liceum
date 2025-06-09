package main

import(
	"fmt"
)

func main7(){
	var pl1, pl2 string
	fmt.Scanln(&pl1)
	fmt.Scanln(&pl2)

	switch {
	case (pl1 == "камень" && pl2 == "ножницы") || (pl1 == "бумага" && pl2 == "камень") || (pl1 == "ножницы" && pl2 == "бумага"):
		fmt.Println("Первый игрок победил")
	case (pl2 == "камень" && pl1 == "ножницы") || (pl2 == "бумага" && pl1 == "камень") || (pl2 == "ножницы" && pl1 == "бумага"):
		fmt.Println("Второй игрок победил")
	default:
		fmt.Println("Ничья")
	}	
}