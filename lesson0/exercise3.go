package main

import "fmt"

func main3(){
	var name string
	var number int
	var password int
	var time float64

	fmt.Scanln(&name)
	fmt.Scanln(&number)
	fmt.Scanln(&password)
	fmt.Scanln(&time)

	fmt.Println(
		fmt.Sprintf(
			"Привет, %s! Приглашаю тебя на соревнование по программированию, которое пройдёт, как всегда, в квартире %d. Оно будет длиться примерно %.1f часа. Не забудь секретный пароль для входа: %d.",
			name,
			number,
			time,
			password,
		),
	)
}