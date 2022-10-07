package main

import "fmt"

func main() {

	number1 := 5
	number2 := 2

	//Sumar
	result := number1 + number2
	fmt.Println(" Suma: ", result)

	//Resta
	result = number1 - number2
	fmt.Println(" Resta: ", result)

	//Multiplicacion
	result = number1 * number2
	fmt.Println(" Multiplicacion ", result)

	//Division
	var splitt float64 = 5.0 / 2.2
	fmt.Println("Division: ", splitt)

	//Modulo

	var modd float64 = 5 % 2
	fmt.Println("Modulo: ", modd)
}
