package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) editName() {

	p.name = "Daniel Melero"
	fmt.Println(p.name)
}

func (p *person) showData() {
	fmt.Println(p.name, " ", p.age)
}

//Herencia
type empleado struct {
	person
	sueldo float32
}

func main() {
	fmt.Println("Hola")
	p := person{"Fabian", 28}
	p.showData()
	p.editName()

	//aplicar herencia
	empl1 := empleado{sueldo: 3500.000} //se muestra con dos llaves
	empl1.name = "Julian"
	empl1.age = 29
	empl1.showData()
	fmt.Println("///////////////////////////////////////////////////////")
	fmt.Println(empl1)
}
