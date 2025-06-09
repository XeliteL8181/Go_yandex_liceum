package main

import(
	"fmt"
)

func main8(){
	var zn string
	var temp int
	fmt.Scanln(&zn, &temp)

	switch{
	case zn == "+" && temp > 20:
		fmt.Println("Стоит надеть майку и шорты")
	case zn == "+" && (temp >= 10 && temp <= 20):
		fmt.Println("Стоит надеть штаны и кофту")
	case zn == "-" && temp > 5:
		fmt.Println("Стоит надеть зимнюю куртку")
	default:
		fmt.Println("Стоит надеть куртку")
	}
}