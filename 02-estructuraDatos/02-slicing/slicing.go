package main

import "fmt"

func main() {
	fmt.Println("Yes!")
	var myslice = []int{4, 5, 6}
	fmt.Println(myslice)
	//add data
	myslice = append(myslice, 7)
	fmt.Println(myslice)

	submyslice := myslice[:2]

	myslice[1] = 55

	fmt.Println(submyslice)
	fmt.Println(myslice)

	//Punteros

	//Longitud

	//Capacidad

	mounths := []string{"Enero", "Febrero", "Marzo"}
	fmt.Printf("Longitud %v , capacidad %v  , espacio memoria %p \n", len(mounths), cap(mounths), mounths)
	mounths = append(mounths, "Abril")
	fmt.Printf("Longitud %v , capacidad %v  , espacio memoria %p \n", len(mounths), cap(mounths), mounths)
}
