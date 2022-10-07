package main

import "fmt"

func pluss(name string, numbers ...int) (string, int) {
	var result int
	message := fmt.Sprintf("La suma de %s es: ", name)
	for _, value := range numbers {
		result += value
	}
	return message, result
}

func main() {
	msg, result := pluss("Fabian", 5, 5, 1, 0, 8, 2, 0, 20, 100)
	fmt.Println(msg, result)
}
