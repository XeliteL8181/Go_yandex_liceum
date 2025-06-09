package main

import (
	"fmt"
	"time"
)

func main9(){
	var now string
	fmt.Scanln(&now)
	now1, _ := time.Parse("2006-01-02/15:04:05", now)
	fmt.Println(
		fmt.Sprintf("Текущее время %d часов, %d минут. Ты точно не забыл про важные дела на сегодня?",
			now1.Hour(),
			now1.Minute(),
		),
	)
}