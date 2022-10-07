package main

import "fmt"

func main() {
	//array
	var arraya1 [3]string
	arraya1[0] = "Daniel"
	arraya1[1] = "Daniel"
	arraya1[2] = "Daniel"

	fmt.Println(arraya1)

	arraya2 := [3]string{"Cabron", "Cabron", "Cabron"}
	fmt.Println(arraya2)

	var slicing1 []int

	for i := 0; i < 10; i++ {
		slicing1 = append(slicing1, i)
		fmt.Println(slicing1)
	}

	var phone [3]string
	phone[0] = "White"
	phone[1] = "Alambrico"
	phone[2] = "Motorola"

	fmt.Println("characteristics phone :", phone)

	display := [2]string{"AOC", "19"}
	fmt.Println("Display : ", display)

	furnitures := [8]string{"Escritorio", "Mueble escritorio", "Pantallas", "Telefono", "Silla visitantes", "Silla empleados", "Iluminacion", "Decoracion"}
	var studyRoom []string
	for _, furniture := range furnitures {
		studyRoom = append(studyRoom, furniture)
	}
	fmt.Println(furnitures)

}
