package main

import "fmt"

func main() {
	var number1 int
	fmt.Println("Digite un numero entero ")
	fmt.Scanln(&number1)

	fmt.Println(" tu numero es:  ", number1)

	for i := 1; i <= number1; i++ {

		if i == 1 {
			fmt.Println(" El numero es 1 ")
			continue
		}
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
