package main

import "fmt"

func main0() {
    fmt.Println("Начало")
    panic("Что-то пошло не так")
    fmt.Println("Конец")
}