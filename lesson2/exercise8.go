package main

import (
	"errors"
	"fmt"
)

var (
	err1 = errors.New("Ошибка: некорректный ввод")
	err2 = errors.New("Ошибка: невалидный возраст")
	err3 = errors.New("Ошибка: невалидное имя")
	err4 = errors.New("Ошибка: невалидная серия и номер паспорта")
)

func main7() {
	var age int
	var name, passportSeriesAndNuber string
	_, err := fmt.Scanln(&age, &name, &passportSeriesAndNuber)
	if err != nil {
		fmt.Println(err1)
		return
	}

	if age < 14 || age > 150 {
		fmt.Println(err2)
		return
	}

	if len(name) < 2 {
		fmt.Println(err3)
		return
	}

	if len(passportSeriesAndNuber) != 10 {
		fmt.Println(err4)
		return
	}

	defer fmt.Println(
		fmt.Sprintf(
			"Имя: %s. Возраст: %d. Серия и номер паспорта: %s",
			name, age, passportSeriesAndNuber,
		),
	)
}
