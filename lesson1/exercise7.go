package main

import (
	"fmt"
)

func main6() {
	var pas1, pas2 string
	fmt.Scanln(&pas1)
	fmt.Scanln(&pas2)

	if len(pas1) >= 8 && len(pas2) >= 8 {
		fmt.Println("Оба пароля надёжные")
		return
	} else if len(pas1) >= 8 && len(pas2) < 8 {
		fmt.Println("Только первый пароль надёжный")
		return
	} else if len(pas1) < 8 && len(pas2) >= 8 {
		fmt.Println("Только второй пароль надёжный")
		return
	}
	fmt.Println("Оба пароля ненадёжные")
}
