package main

import (
	"strings"
)

func NumbersToLetters(input string) string {
	otv := strings.Replace(input, "0", "ноль", -1)
	otv = strings.Replace(otv, "1", "один", -1)
	otv = strings.Replace(otv, "2", "два", -1)
	otv = strings.Replace(otv, "3", "три", -1)
	otv = strings.Replace(otv, "4", "четыре", -1)
	otv = strings.Replace(otv, "5", "пять", -1)
	otv = strings.Replace(otv, "6", "шесть", -1)
	otv = strings.Replace(otv, "7", "семь", -1)
	otv = strings.Replace(otv, "8", "восемь", -1)
	otv = strings.Replace(otv, "9", "девять", -1)
	otv = strings.Replace(otv, "+", "плюс", -1)
	otv = strings.Replace(otv, "-", "минус", -1)
	otv = strings.Replace(otv, "*", "умножить на", -1)
	otv = strings.Replace(otv, "/", "разделить на", -1)
	return otv
}