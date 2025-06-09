package main

import ( 
	"fmt"
	"time"
)

func main4(){
	now0 := time.Now()
	fmt.Println("Current time: ", now0)

	timeOfBirth := time.Date(2006, 06, 14, 23, 32, 57, 0, time.UTC)
	diff := now0.Sub(timeOfBirth) // Разница между текущим и временем рождения
	fmt.Println(diff.Hours()) // Вывод разницы в часах


	now := time.Now() // Примеры форматирования времени
	fmt.Println("1. Время в формате RFC3339:", now.Format(time.RFC3339))
	fmt.Println("2. Полная дата и время:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("3. Краткая дата:", now.Format("2006-01-02"))
	fmt.Println("4. Время в 24-часовом формате:", now.Format("15:04:05"))
	fmt.Println("5. Время в 12-часовом формате с AM/PM:", now.Format("3:04 PM"))
	fmt.Println("6. День недели:", now.Format("Monday"))
	fmt.Println("7. Сокращённый месяц:", now.Format("January"))
}