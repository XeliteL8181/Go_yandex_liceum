package main

import "fmt"

func main0() {
	var name string
	fmt.Println("Введи своё имя:") 
	fmt.Scanln(&name)

	var age int
	fmt.Println("Введи свой возраст: ")
	fmt.Scanln(&age)
	
	var mark float64
	fmt.Println("Введи свой средний балл по любимому предмету: ")
	fmt.Scanln(&mark)

	const secretNumber = 5252

	fmt.Println(
		fmt.Sprintf(
			"Привет, %d-летний уёбок под именем %s. Я знаю, что ты ебанный пубертатник за всю свою жизнь подрочил на %d хентачей, ПОГОДИ ЧТО, НИХУЯСЕ, СТОЛЬКО ВООБЩЕ СУЩЕСТВУЕТ? ТЕПЕРЬ ПОНЯТНО, ПОЧЕМУ ДЕВУШКИ В СРЕДНЕМ ОЦЕНИВАЮТ ТЕБЯ НА %f. Бай, ебанный дрочер",
			age, 
			name, 
			secretNumber,
			mark,
		),
	)
}