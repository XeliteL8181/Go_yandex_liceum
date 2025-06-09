package main

import (
	"fmt"
	"math"
)

func main8(){
	var a, b float64
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println(math.Round(math.Max(a, b)))
}