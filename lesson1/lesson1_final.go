package main

import (
	"fmt"
)

func main() {
	var name string
	var num, queue int
	st1, st2, st3, st4, st5 := "-", "-", "-", "-", "-"
	fmt.Scanln(&name, &num)

	for name != "конец" {

		if name == "" {
			name, num = "", 0
			fmt.Scanln(&name, &num)
			continue
		}

		if name == "количество" {
			fmt.Println(
				fmt.Sprintf(
					"Осталось свободных мест: %d\nВсего человек в очереди: %d",
					5-queue, queue,
				),
			)
		} else if name == "очередь" {
			fmt.Println(
				fmt.Sprintf(
					"1. %s\n2. %s\n3. %s\n4. %s\n5. %s",
					st1, st2, st3, st4, st5,
				),
			)
		} else {

			switch {
			case num < 1 || num > 5:
				fmt.Println(
					fmt.Sprintf(
						"Запись на место номер %d невозможна: некорректный ввод",
						num,
					),
				)

			case (st1 != "-" && st2 != "-" && st3 != "-" && st4 != "-" && st5 != "-"):
				fmt.Println(
					fmt.Sprintf("Запись на место номер %d невозможна: очередь переполнена", num),
				)

			case num == 1:
				if st1 == "-" {
					st1 = name
				} else {
					fmt.Println(
						fmt.Sprintf(
							"Запись на место номер %d невозможна: место уже занято",
							num,
						),
					)
				}
				queue += 1

			case num == 2:
				if st2 == "-" {
					st2 = name
				} else {
					fmt.Println(
						fmt.Sprintf(
							"Запись на место номер %d невозможна: место уже занято",
							num,
						),
					)
				}
				queue += 1

			case num == 3:
				if st3 == "-" {
					st3 = name
				} else {
					fmt.Println(
						fmt.Sprintf(
							"Запись на место номер %d невозможна: место уже занято",
							num,
						),
					)
				}
				queue += 1

			case num == 4:
				if st4 == "-" {
					st4 = name
				} else {
					fmt.Println(
						fmt.Sprintf(
							"Запись на место номер %d невозможна: место уже занято",
							num,
						),
					)
				}
				queue += 1

			case num == 5:
				if st5 == "-" {
					st5 = name
				} else {
					fmt.Println(
						fmt.Sprintf(
							"Запись на место номер %d невозможна: место уже занято",
							num,
						),
					)
				}
				queue += 1
			}
		}

		name, num = "", 0
		fmt.Scanln(&name, &num)
	}

	fmt.Println(
		fmt.Sprintf(
			"1. %s\n2. %s\n3. %s\n4. %s\n5. %s",
			st1, st2, st3, st4, st5,
		),
	)
}
