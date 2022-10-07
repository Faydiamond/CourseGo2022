package main

import "fmt"

type heroe struct {
	Name string
	Type string
	Age  int
}

func main() {
	/*////////////////////////////////////////////////////////////////////////////////////////////////////*/
	//arrays
	var names [6]string
	names[0] = "Fabian"
	names[1] = "Andres"
	names[2] = "Julian"
	names[3] = "Catalina"
	names[4] = "Monica"
	names[5] = "Estefania"
	fmt.Println(names)
	for _, name := range names {
		fmt.Println(name)
	}

	lastName := [6]string{"Barbon", "Estrada", "Perez", "Sanchez", "Gutierrez", "Lopez"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i], " ", lastName[i])
	}
	/*////////////////////////////////////////////////////////////////////////////////////////////////////*/
	/*////////////////////////////////////////////////////////////////////////////////////////////////////*/
	//slicen
	var boocks []string

	boocks = append(boocks, "El diario de Ana Frank")
	boocks = append(boocks, "Primero es la salud tomo #1")
	boocks = append(boocks, "Primero es la salud tomo #3")
	boocks = append(boocks, "Primero es la salud tomo #2")
	boocks = append(boocks, "Algebra moderna")
	boocks = append(boocks, "Computacion 1")
	boocks = append(boocks, "paracaidas y vueltas")
	boocks = append(boocks, "El principito")
	boocks = append(boocks, "El mundo de sofia")

	for _, boock := range boocks {
		if boock == "El principito" {
			fmt.Println("Es un buen libro apto para todo publico.")
		}
	}
	/*////////////////////////////////////////////////////////////////////////////////////////////////////*/
	/*////////////////////////////////////////////////////////////////////////////////////////////////////*/
	//mapa
	carBuy := make(map[string]float32, 8)
	carBuy["leche"] = 3.550
	carBuy["huevos"] = 550
	carBuy["mantequilla"] = 2.550
	carBuy["yogurt"] = 13.250
	carBuy["galletas tosh"] = 5.780
	carBuy["mani"] = 1.550
	carBuy["bolsa para la basura"] = 650.4
	carBuy["filete de salmon"] = 63.710

	for _, itemCar := range carBuy {
		fmt.Println(itemCar)
	}
	/*////////////////////////////////////////////////////////////////////////////////////////////////////*/
	/*////////////////////////////////////////////////////////////////////////////////////////////////////*/
	//struct
	t := heroe{Name: "Thor", Age: 29, Type: "Electrico"}
	a := heroe{Name: "Antman", Age: 35, Type: "Insecto"}

	fmt.Println(t)
	fmt.Println(a)
}
