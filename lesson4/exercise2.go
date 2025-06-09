package main

import (
	"fmt"
)

type Employee struct {
	name string
	position string
	salary float64
	bonus float64
}

func (empl Employee) CalculateTotalSalary() {
	fmt.Println("Employee:", empl.name) 
	fmt.Println("Position:", empl.position)
	fmt.Printf("Total Salary: %.2f", empl.salary + empl.bonus)
}