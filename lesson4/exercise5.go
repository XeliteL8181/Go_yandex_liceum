package main

import ()

type Animal1 interface {
	MakeSound() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (dog Dog) MakeSound() string {
	return "Гав!"
}

func (cat Cat) MakeSound() string {
	return "Мяу!"
}

func SoundOfAnimal(animal Animal1) string {
	return animal.MakeSound()
}