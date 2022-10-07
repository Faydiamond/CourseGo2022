package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		error := recover()
		if error != nil {
			fmt.Println("Ups se ha manejado el panico y el programa se ha ejecutado de manera correcta")
		}
	}()

	if file, error := os.Open("text56.txt"); error != nil {
		panic("No es posible obtener el contenido del archivo!")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo")
			file.Close()
		}()
		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println("Contenido archivo:  ", contenidoArchivo)
	}
}
