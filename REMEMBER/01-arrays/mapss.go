package main

import "fmt"

func main() {
	var carBuy = make(map[string]int)
	carBuy["leche"] = 12
	carBuy["huevos"] = 6
	carBuy["mantequilla"] = 2

	fmt.Println(carBuy)

	for i, item := range carBuy {
		fmt.Println(i, item)
	}

}
