package main

import "fmt"

func main() {
	//solo es visible dentro del if
	if name, age := "Daniel", 52; name == "Harry" {
		fmt.Println("Hola, ", name)
	} else {
		fmt.Printf(" %s con %d anos \n", name, age)
	}

	var students = make(map[string]string)
	students["Daniel"] = "Daniel@Mail.com"
	students["Pablo"] = "Pablo@Mail.com"
	students["Maycol"] = "Maycol@Mail.com"
	students["Julieta"] = "Julieta@Mail.com"

	correo, ok := students["Julieta"]
	fmt.Println(correo, ok)

	if correo, ok := students["chavito"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println(" No existe el usuario ingresado ")
	}

}
