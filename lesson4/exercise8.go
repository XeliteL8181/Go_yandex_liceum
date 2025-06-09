package main

import (
	"fmt"
)

type User1 struct {
	Name string
	Age int
	IsActive bool
}

func NewUser(name string, age int) (*User1, error) {
	if (name == "") {
		return nil, fmt.Errorf("name is empty for user")
	}

	if age == 0 {
		age = 18
	}

	return &User1{Name: name, Age: age, IsActive: true}, nil
}