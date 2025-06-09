package main

import(
	"fmt"
)

func main0(){
	var Go string;
	fmt.Scanln(&Go)
	if (Go == "Go"){
		fmt.Println("Go!")
		return
	}
	fmt.Println("Я знаю только Go!")
}