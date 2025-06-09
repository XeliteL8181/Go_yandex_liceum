package main

import (
	"fmt"
	"lesson/students"
)

func main1() {
	student1 := students.Student{}
	fmt.Println(student1)

	user1 := students.User{ Gender: "male", Name: "Artyom", Email: "fuckYou@gmail.com" }
	fmt.Println(user1)
}