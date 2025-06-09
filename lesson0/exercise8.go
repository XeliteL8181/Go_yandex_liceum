package main

import (
	"fmt"
	"time"
)

func main10(){
	var futureS, pastS string
	fmt.Scanln(&futureS)
	fmt.Scanln(&pastS)
	
	future, _ := time.Parse("2006-01-02", futureS)
	past, _ := time.Parse("2006-01-02", pastS)

	diff1 := future.Year() - past.Year()
	fmt.Println(
        fmt.Sprintf("%d year ago", diff1),
    )
}