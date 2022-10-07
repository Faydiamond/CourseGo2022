package main

import "fmt"

func main() {
	var names = []string{"Daniel", "Julieta", "Julian", "Carlos"}
	fmt.Println(names)
	subnames := names[2:]
	fmt.Println(subnames)
}
