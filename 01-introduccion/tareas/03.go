package main

import "fmt"

func igv(a float32) float32 {
	apigv := a * 18 / 100
	total := a + apigv
	return total
}

func main() {
	fmt.Println(" Ingrese el valor del producto ")
	var number1 float32
	fmt.Scanln(&number1)
	fmt.Println(" El valor del producto es: ", number1)
	fmt.Println("El valor del igv es del 18%")
	fmt.Println("El precio de venta es: ", igv(100))
}
