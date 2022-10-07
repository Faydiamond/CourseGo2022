package main

import "fmt"

var mensaje string

func saludo1() {
	mensaje = "Hola desde la funcion1"
	fmt.Println(mensaje)
}

func saludo2() {
	mensaje = "Hola desde la funcion2"
	fmt.Println(mensaje)
}

func main() {
	mensaje = "Hola desde la funcion principal"
	fmt.Println(mensaje)
	defer saludo1() //se ejecuta al final
	saludo2()
	fmt.Println(mensaje)
}
