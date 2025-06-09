package main

import (
	"fmt"
)

func PrintFlightRow(num, start, final string, time float64, table, gate int, registration bool) {
	if registration == false {
		fmt.Println(
			fmt.Sprintf(
				"| %s | %s—%s | %d регистрация продолжается |",
				num, start, final, table,
			),
		)
	} else {
		fmt.Println(
			fmt.Sprintf(
				"| %s | %s—%s | регистрация закончилась, проходите к гейту: %d | длительность полёта %.1f часа |",
				num, start, final, gate, time,
			),
		)
	}
}
