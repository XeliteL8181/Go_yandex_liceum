package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
	address string
}

func (person Person) PrettyPrint() {
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)
	fmt.Println("Address:", person.address)
}