package main

import "fmt"

func main() {
	a := 10
	b := 20

	var r bool
	//iguales
	r = a == b
	fmt.Printf("%d es igual a %d? %t ", a, b, r)

	//distinto
	r = a != b
	fmt.Printf("%d es distinto a %d? %t ", a, b, r)
}
