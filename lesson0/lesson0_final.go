package main

import (
	"fmt"
	"time"
)

func main() {
	var dateS, name, lastName, fatherName string
	var vip1, vip2, vip3 float64
	fmt.Scanln(&dateS)
	fmt.Scanln(&name)
	fmt.Scanln(&lastName)
	fmt.Scanln(&fatherName)
	fmt.Scanln(&vip1)
	fmt.Scanln(&vip2)
	fmt.Scanln(&vip3)

	date, _ := time.Parse("02.01.2006", dateS)
	dateI := date.AddDate(0, 0, 15)
	vipRub := int(vip1 + vip2 + vip3)
	vipCop := int((vip1 + vip2 + vip3) * 100) % 100

	fmt.Println(
		fmt.Sprintf(
			"Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по выполненной вами работы. \nДата подписания договора: %.2d.%.2d.%.4d.",
			lastName, name, fatherName, dateI.Day(), dateI.Month(), dateI.Year(),
		),
		fmt.Sprintf(
			"Просим вас подойти в офис в любое удобное для вас время в этот день. \nОбщая сумма выплат составит %d руб. %d коп.",
			vipRub, vipCop,
		),
		fmt.Sprintf(
			"\n\nС уважением, \nГл. бух. Иванов А.Е.",
		),
	)
}