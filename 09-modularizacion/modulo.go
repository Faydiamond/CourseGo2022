package main

import (
	"fmt"
	"paquetes/figuras"
	"paquetes/mensajes"
	"paquetes/models"
)

func main() {
	mensajes.Hola()
	mensajes.Imprimir()

	cua01 := figuras.Cuadrado{5, 5}
	figuras.CalcularValor(&cua01)

	person1 := models.Person{}
	person1.Construct("Daniel ", 28)
	fmt.Println(person1)

	fmt.Println(person1.GetNombre())
	person1.SetNombre("Mariano")
	fmt.Println(person1)
}
