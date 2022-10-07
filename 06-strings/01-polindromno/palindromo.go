package main

import (
	"fmt"
	"strings"
)

func reverse(mytext string) string {
	arMytext := strings.Split(mytext, " ")
	arnvert := make([]string, 0)

	for i := len(arMytext) - 1; i >= 0; i-- {
		arnvert = append(arnvert, arMytext[i])
	}

	return strings.Join(arnvert, "")
}

func isPalindromo(worldd string) bool {
	worldd = strings.ToLower(worldd)
	worldd = strings.ReplaceAll(worldd, "z", "s")

	invertWorld := reverse(worldd)
	return invertWorld == worldd
}

func main() {
	if isPalindromo("ana") {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}

}
