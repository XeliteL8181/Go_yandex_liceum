package main

import (
	"fmt"
)

type Animal interface {
	MakeSound() string 
	GetName() string
	GetInfo() string
}

type animal struct {
	name string
	species string
	age int
	sound string
}

type ZooKeeper struct {}

func (a animal) MakeSound() string {
	return a.sound
}

func (a animal) GetName() string {
	return a.name
}

func (a animal) GetInfo() string {
	return fmt.Sprintf(
		"Имя: %s, Вид: %s, Возраст: %d",
		a.name, a.species, a.age,
	)
}

func NewAnimal(name, species string, age int, sound string) Animal {
	return &animal{ name: name, species: species, age: age, sound: sound }
}

func ZooShow(animals []Animal) {
	for _, animal := range animals {
		fmt.Println(animal.GetInfo())
		fmt.Println(animal.MakeSound())
	}
}

func (z ZooKeeper) Feed(animal Animal) {
	name := animal.GetName()
	sound := animal.MakeSound()
	
	fmt.Println(
		fmt.Sprintf(
			"Смотритель зоопарка кормит %s. %s!",
			name, sound,
		),
	)
}