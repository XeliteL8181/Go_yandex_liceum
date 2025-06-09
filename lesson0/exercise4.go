package main

import (
	"fmt"
	"math"
)

func main6(){
	var ars, gosha, ira float64
	fmt.Scanln(&ars)
	fmt.Scanln(&gosha)
	fmt.Scanln(&ira)
	fmt.Println(math.Min(math.Min(ars, gosha), ira))
}