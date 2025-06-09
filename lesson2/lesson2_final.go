package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var(
	errInput = errors.New("исправь свой ответ, а лучше ложись поспать")
)

func currentDayOfTheWeek() string {
	nowWD := time.Now().Weekday().String()
	
	switch nowWD{
	case "Monday":
		return "Понедельник"
	case "Tuesday":
		return "Вторник"
	case "Wednesday":
		return "Среда"
	case "Thursday":
		return "Четверг"
	case "Friday":
		return "Пятница"
	case "Saturday":
		return "Суббота"
	default:
		return "Воскресенье"
	}
}

func dayOrNight() string {
	nowDN := time.Now().Hour()
	switch{
	case nowDN >= 10 && nowDN <= 22:
		return "День"
	default:
		return "Ночь"
	} 
}

func nextFriday() int {
	nowWD := time.Now().Weekday().String()
	
	switch nowWD{
	case "Monday":
		return 4
	case "Tuesday":
		return 3
	case "Wednesday":
		return 2
	case "Thursday":
		return 1
	case "Friday":
		return 0
	case "Saturday":
		return 6
	default:
		return 5
	}
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	WD := currentDayOfTheWeek()
	if strings.ToLower(answer) == strings.ToLower(WD) {
		return true
	}
	return false
}

func CheckNowDayOrNight(answer string) (bool, error) {
	DN := dayOrNight()
	if strings.ToLower(answer) == "день" || strings.ToLower(answer) == "ночь"{
		if strings.ToLower(answer) == strings.ToLower(DN) {
			return true, nil
		} else if strings.ToLower(answer) != strings.ToLower(DN){
			return false, nil
		}
	}
	return false, errInput
}

func main() {
	fmt.Println(currentDayOfTheWeek())
	fmt.Println(dayOrNight())
	fmt.Println(nextFriday())
	fmt.Println(CheckCurrentDayOfTheWeek("ПонеДЕльник"))
	fmt.Println(CheckNowDayOrNight("НоЧЬ"))
}
