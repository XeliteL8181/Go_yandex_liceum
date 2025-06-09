package main

import(
	"fmt"
	"math"
)

func main3(){
	var num float64
	fmt.Scanln(&num)

	if (math.Pow(math.Sqrt(num), 2) == num){
		fmt.Println(math.Sqrt(num))
		return
	}
	fmt.Println(-1)
}