package main

import (
	"fmt"
	"math"
)

func main5() {
	fmt.Println(math.Pi)             // Число пи
	fmt.Println(math.E)              // Экспонента
	fmt.Println(math.Sqrt2)          // Корень 2
	fmt.Println(math.Abs(-1.0))      // Возвращает модуль числа
	fmt.Println(math.Min(1.0, -1.0)) // Возвращает минимальное число
	fmt.Println(math.Max(-1.0, 1.0)) // Возваращает максимальное число
	fmt.Println(math.Sqrt(2.0))      // Возваращает корень числа
	fmt.Println(math.Sin(45.0))      // Возвращает значение синуса угла
	fmt.Println(math.Pow(2.0, 3.0))  // Возводит в степень
	fmt.Println(math.Ceil(5.2))      // Округляет к большему
	fmt.Println(math.Floor(5.2))     // Округляет к меньшему
	fmt.Println(math.Round(5.2))     // Округляет к ближайшему
}