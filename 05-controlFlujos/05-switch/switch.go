package main

import "fmt"

func main() {
	var vocal string
	fmt.Println("Ingrese una vocal")
	fmt.Scanln(&vocal)
	/*
		switch {
		case vocal == "a" || vocal == "A":
			fmt.Println("Has ingresado la vocal a")
		case vocal == "e" || vocal == "E":
			fmt.Println("Has ingresado la vocal e")
		case vocal == "i" || vocal == "I":
			fmt.Println("Has ingresado la vocal i")
		case vocal == "o" || vocal == "O":
			fmt.Println("Has ingresado la vocal o")
		case vocal == "u" || vocal == "U":
			fmt.Println("Has ingresado la vocal u")
		default:
			fmt.Println(vocal, " no es una vocal como tal.")
		}*/

	switch vocal {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
		fmt.Println(vocal, "Es una vocal")
	default:
		fmt.Println(vocal, " no es una vocal como tal.")
	}
}
