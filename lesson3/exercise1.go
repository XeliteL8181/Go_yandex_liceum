package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func CountLengthAndBytes(first, second string) string {
	words := []string{first, second}
	firstEx := strings.Join(words, "")
	secondEx := len(firstEx)
	thirdEx := utf8.RuneCountInString(firstEx)
	return fmt.Sprintf(
		"Объединённая строка: %s. Количество байт: %d. Количество символов: %d.", 
		firstEx, secondEx, thirdEx,
	)
}
