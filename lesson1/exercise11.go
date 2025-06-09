package main

import(
	"fmt"
)

func main10(){
	var n int
	var mark float64
	fmt.Scanln(&n)

	for i := 0; i < n; i++{
		fmt.Scanln(&mark)
		switch{
		case mark >= 0 && mark <= 49:
			fmt.Println(2)
		case mark >= 50 && mark <= 74:
			fmt.Println(3)
		case mark >= 75 && mark <= 89:
			fmt.Println(4)
		case mark >= 90 && mark <= 100:
			fmt.Println(5)
		default:
			fmt.Println("Неверный балл")
		}
	}
}