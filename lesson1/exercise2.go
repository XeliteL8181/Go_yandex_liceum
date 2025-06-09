package main

import(
	"fmt"
)

func main1(){
	var num1, num2 int
	fmt.Scanln(&num1, &num2)
	if (num1 > num2){
		fmt.Println("Первое число больше второго")
		return
	} else if (num1 == num2){
		fmt.Println("Числа равны")
		return
	}
	fmt.Println("Второе число больше первого")
}