package main

import "fmt"

func main() {
	func() {
		fmt.Println(" Ejecutar funcion anonima ")
	}()

	myVar := func() {
		fmt.Println(" Ejecutar funcion anonima en una variable")
	}

	myVar()
}
