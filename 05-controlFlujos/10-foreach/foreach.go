package main

import "fmt"

func main() {
	fruits := [...]string{"Pera", "Manzana", "Durazno"}

	for _, value := range fruits {
		fmt.Println(value)
	}
}
