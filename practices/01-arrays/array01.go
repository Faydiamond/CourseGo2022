package main

import "fmt"

func main() {
	var elemetsBooks = [3]string{"Pasta", "Hojas", "cantidad de hojas"}
	for i := 0; i < len(elemetsBooks); i++ {
		fmt.Println(elemetsBooks[i])
	}
	fmt.Println("/////////////////////////////////////////////////////////////")
	car := [...]string{"color", "marca", "modelo", "cilindraje", "transmision"}
	for _, v := range car {
		fmt.Println(v)
	}

	var shampoo [2]string
	shampoo[0] = "Marca"
	shampoo[1] = "peso"

	fmt.Println(shampoo)
}
