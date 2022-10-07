package main

import "fmt"

func main() {
	var number1, number2 int
	fmt.Println(" Ingrese el numero1 ")
	fmt.Scanln(&number1)
	fmt.Println(" Ingrese el numero2 ")
	fmt.Scanln(&number2)

	cocient := number1 / number2
	fmt.Println("El cociente es: ", cocient)
	var resi int
	resi = number1 % number2
	fmt.Println("El cociente es: ", resi)

}
