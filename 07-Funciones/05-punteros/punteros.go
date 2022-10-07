package main

import "fmt"

func modifyText(text *string) {
	fmt.Printf("%p \n", &text)
	*text = "Hola desde la funcion que modifica la cadena "
}

func main() {

	text := "Hola Golang"
	fmt.Printf("%p \n", &text)
	fmt.Println("Antes de la funcion: ", text)
	modifyText(&text)
	fmt.Println("Despues de la funcion: ", text)
}
