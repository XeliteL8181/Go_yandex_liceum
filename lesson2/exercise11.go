package main

import (
	"fmt"
	"strings"
)

func CheckLetters(text string) string {
	otv := strings.Count(text, "е")
	if otv == 0 {
		return "Текст готов к публикации!"
	}
	return fmt.Sprintf("Количество возможных ошибок: %d, перепроверьте текст", otv)
}