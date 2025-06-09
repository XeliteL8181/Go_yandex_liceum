package main

import(
	"fmt"
)

func main2(){
	var rub, cop, coffee int
	fmt.Scanln(&rub, &cop, &coffee)
	
	if ((rub + cop / 100) >= coffee){
		fmt.Println("Сегодня будет вкусный кофе!")
		return
	}
	fmt.Println("Стоит подкопить")
}