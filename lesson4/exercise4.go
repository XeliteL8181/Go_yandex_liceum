package main

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width float64
	height float64
}

func (circ Circle) Area() float64 {
	return math.Pi * circ.radius * circ.radius
}

func  (rect Rectangle) Area() float64 {
	return rect.width * rect.height
}

func CalculateArea(s Shape) float64 {
	return s.Area()
}