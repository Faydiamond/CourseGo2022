package main

import (
	"fmt"
	"strings"
)

func repeat(number int) func(text string) string {
	return func(text string) string {
		msg := fmt.Sprintf("Hola  %s ", text)
		return strings.Repeat(msg, number)
	}
}

func main() {
	repeat5 := repeat(5)
	fmt.Println(repeat5("Cabron "))

}
