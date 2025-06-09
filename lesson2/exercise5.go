package main

import(
	"math"
)

func findDiscriminant(a, b, c float64) float64 {
	return b * b - 4 * a * c
}

func SquareRoots(a, b, c float64) (float64, float64) {
	Discriminant := findDiscriminant(a, b, c)
	
	if (Discriminant > 0){
		Sq1 := ((-b - math.Sqrt(Discriminant)) / (2 * a))
		Sq2 := ((-b + math.Sqrt(Discriminant)) / (2 * a))
		return Sq1, Sq2
	} else if (Discriminant == 0) {
		Sq := (-b / (2 * a))
		return Sq, Sq
	}
	return 0, 0
}