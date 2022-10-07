package main

import "fmt"

func main() {
	days := make(map[int]string)
	fmt.Println("Dias: ", days)
	//clave valor
	days[0] = "Lunes"
	days[1] = "Martes"
	days[2] = "Miercoles"
	days[3] = "Jueves"
	days[4] = "Viernes"
	days[5] = "Sabado"
	days[6] = "Domingo"
	fmt.Println("Dias: ", days)
	//modificar
	days[2] = "MIERCOLES"
	fmt.Println("Dias: ", days)
	//eliminar
	delete(days, 6)
	fmt.Println("Dias: ", days)

	students := make(map[string][]int)
	students["fabian"] = []int{5, 4, 3}
	students["daniel"] = []int{5, 4, 3}
	fmt.Println(students)
	fmt.Println(students["daniel"])
	fmt.Println(students["daniel"][2])
}
