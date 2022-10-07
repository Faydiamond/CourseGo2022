package main

import "fmt"

func gretting(name string) {
	fmt.Println(" Hola, ", name)
}

func plus(x, y int) int {
	return x + y
}

func main() {
	gretting("Fabian")
	result := plus(1250, 2025)
	fmt.Println(result)
}
