package main

import (
	"fmt"
)

func main() {
	//arreglos
	var controlXbox [4]string
	controlXbox[0] = "tamano"
	controlXbox[1] = "generico"
	controlXbox[2] = "color"
	controlXbox[3] = "alambrico"

	//mostrar datos array
	for i, controlItem := range controlXbox {
		fmt.Println(i, controlItem)
	}

	library := [4]string{"material", "color", "tamano", "divisiones"}
	fmt.Println("Libreria ", library)

	//slicen
	var garden []string
	garden = append(garden, "hortensias")
	garden = append(garden, "novios")
	garden = append(garden, "geranios")
	garden = append(garden, "geranios enanos")
	garden = append(garden, "buganvillas")
	fmt.Println(garden)

	var bag []int

	for i := 0; i <= 25; i++ {
		bag = append(bag, i)
		fmt.Println(bag)

	}

	certificate := [2]string{"universidad", "ano"}
	fmt.Println("Certificado : ", certificate)

	var key []string

	for j := 0; j < 6; j++ {
		if j == 0 {
			key = append(key, "Color")
		} else if j == 1 {
			key = append(key, "Marca")
		} else if j == 2 {
			key = append(key, "Tamano")
		} else if j == 3 {
			key = append(key, "Material")
		} else if j == 4 {
			key = append(key, "Controles")
		} else if j == 5 {
			key = append(key, "tipo")
		}
		fmt.Println(key[j])
	}

}
