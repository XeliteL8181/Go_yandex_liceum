package main

import(
	"fmt"
)

func myFunc() int{
	a := 1
	return a
}

func main3(){
	a := myFunc()
	fmt.Println(a)
}