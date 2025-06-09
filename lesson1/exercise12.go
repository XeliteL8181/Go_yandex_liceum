package main

import (
	"fmt"
)

func main11(){
	var cnt int
	var price, otv, sale float64
    fmt.Scanln(&cnt)
	fmt.Scanln(&sale)
    sale /= 100

	for i := 0; i < cnt; i++{
		fmt.Scanln(&price)
		otv += price - price * sale
	}
    
	fmt.Println(otv)
}