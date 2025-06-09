package main

import(
	"fmt"
)

func main5(){
	var num1, num2, num3 float64
	fmt.Scanln(&num1, &num2, &num3)
	if (num1 == num2) && (num1 == num3) && (num2 == num3){
		fmt.Println("Максимальное равенство")
		return
	}
	fmt.Println("Не равны")
}