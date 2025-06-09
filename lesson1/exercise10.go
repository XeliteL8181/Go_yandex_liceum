package main

import (
	"fmt"
)

func main9(){
	var Go string
	for i := 0; i < 5; i++{
		fmt.Scanln(&Go)
		if Go == "Go"{
			fmt.Println("Go!")
			continue
		}
		fmt.Println("Я знаю только Go!")
	}
}