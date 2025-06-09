package main

import(
	"fmt"
)

func BuyFries(size string){
	switch{
	case size == "S":
		fmt.Println("Картошка фри будет стоить 49 рублей")
	case size == "M":
		fmt.Println("Картошка фри будет стоить 89 рублей")
	case size == "L":
		fmt.Println("Картошка фри будет стоить 99 рублей")
	default:
		fmt.Println("Некорректный размер")
	}
}

func BuyCola(size string){
	switch{
	case size == "S":
		fmt.Println("Кола будет стоить 99 рублей")
	case size == "M":
		fmt.Println("Кола будет стоить 119 рублей")
	case size == "L":
		fmt.Println("Кола будет стоить 139 рублей")
	default:
		fmt.Println("Некорректный размер")
	}
}

func printPrice(price int, product string){
	fmt.Println(
        fmt.Sprintf(
            "%s будет стоить %d рублей",
            product, price,
        ),
    )
}