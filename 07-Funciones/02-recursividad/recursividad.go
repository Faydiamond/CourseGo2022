package main

import "fmt"

func recursivity(number int) int {
	if number == 0 {
		return 1
	}

	result := number * recursivity(number-1)
	return result
}

func main() {
	factorial := recursivity(5)
	fmt.Println(factorial)
}
