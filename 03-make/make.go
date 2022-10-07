package main

import "fmt"

func main() {
	numbers := make([]int, 3, 3) //longitud, capacidad
	fmt.Println(numbers, len(numbers), cap(numbers))
	numbers[0] = 100
	numbers[1] = 1000
	numbers[2] = 10000
	fmt.Println(numbers, len(numbers), cap(numbers))
	numbers = append(numbers, 100000)
	fmt.Println(numbers, len(numbers), cap(numbers))
}
