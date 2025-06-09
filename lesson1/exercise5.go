package main

import(
	"fmt"
	"math"
)

func main4(){
	var num int
	fmt.Scanln(&num)

	if (num == 0){
		fmt.Println("Число 0")
		return
	}
	if (0 <= math.Abs(float64(num)) && math.Abs(float64(num)) <= 9){
		fmt.Println("Число однозначное")
		return
	}
	if (num % 2 == 0){
		fmt.Println("Число чётное")
		return
	}
	if (num > 0){
		fmt.Println("Число положительное")
		return
	}
	fmt.Println("Число красивое")
}