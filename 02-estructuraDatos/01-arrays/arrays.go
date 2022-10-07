package main

import "fmt"

func main() {
	var one [4]int
	fmt.Println(one)
	one[0] = 5
	one[1] = 15
	one[2] = 25
	one[3] = 35
	fmt.Println(one)
	fmt.Println(one[2])

	names := [2]string{"Daniel", "Melero"}
	fmt.Println(names)

	fruits := [...]string{"Pera", "Manzana", "Naranja"}
	fmt.Println(fruits, len(fruits))
	//indices carater vacio
	coins := [...]string{
		0: "Dolares",
		3: "Euros",
		5: "Dolares Canadienses",
	}
	fmt.Println(coins, len(coins))
	subCoins := coins[:4]
	fmt.Println(subCoins, len(subCoins))
	fmt.Println(coins[4:])
}
