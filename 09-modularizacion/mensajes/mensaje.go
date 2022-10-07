package mensajes

import "fmt"

func Hola() {
	fmt.Println("Hola desde el paquete!")
}

const mensaje = "Soy privado"

func funcionprivada() {
	fmt.Println("Soy una funcion privada")
}

func Imprimir() {
	fmt.Println(mensaje)
	funcionprivada()

}
