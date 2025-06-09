package main

import ()

type Student struct {
	name string
	solvedProblems int
	scoreForOneTask float64
	passingScore float64
}

func (student Student) IsExcellentStudent() bool {
	totalScore := float64(student.solvedProblems) * student.scoreForOneTask
	if (totalScore >= student.passingScore) {
		return true
	} 
	return false
}