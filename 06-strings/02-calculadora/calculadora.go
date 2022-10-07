package main

import (
	"fmt"
	"strconv"
	"strings"
)

func pluss(expres string) int {

	arMytext := strings.Split(expres, "+")
	var result int
	for i := 0; i <= len(arMytext)-1; i++ {
		num, error := strconv.Atoi(arMytext[i])
		if error != nil {
			fmt.Println(" El erro es: ", error)
			fmt.Println("Por favor digite numeros enteros, esta expresion solo acpeta el simbolo +")
		} else {
			result += num
		}
	}

	//fmt.Printf("%T ", arMytext) //tipo de dato
	//fmt.Println("La suma es: ", result)
	return result
}

func main() {
	var expression string
	fmt.Println("Digite su expresion ")
	fmt.Scanln(&expression)
	fmt.Println(pluss(expression))
}
