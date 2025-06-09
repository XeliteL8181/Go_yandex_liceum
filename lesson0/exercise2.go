package main

import "fmt"

func main2(){
	var love string;
	fmt.Scanln(&love)
	fmt.Println(
		fmt.Sprintf("А ещё я люблю %s!", love),
	)
}