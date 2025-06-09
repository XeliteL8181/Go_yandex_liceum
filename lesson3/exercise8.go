package main

import (
	"fmt"
)

func PrettyArrayOutput(array [9]string) {
	for i := 0; i < len(array); i++ {
		if i < 7 {
			fmt.Printf(
				fmt.Sprintf(
					"%d я уже сделал: %s\n",
					i + 1, array[i],
				),
			)
		} else {
			fmt.Printf(
				fmt.Sprintf(
					"%d не успел сделать: %s\n",
					i + 1, array[i],
				),
			)
		}
	}
}