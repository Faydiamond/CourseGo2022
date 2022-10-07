package main

import "fmt"

func pluss(a int, b int) int {
	return a + b
}

func main() {

	fmt.Println("Ingrese el numero 1 ")
	var number1, number2 int
	fmt.Scanln(&number1)
	fmt.Println("Ingrese el numero 2 ")
	fmt.Scanln(&number2)
	fmt.Println("El resultado de sumar los dos numeros es: ", pluss(number1, number2))
}
