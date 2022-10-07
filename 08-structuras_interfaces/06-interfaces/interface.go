package main

import "fmt"

type Animal interface {
	mover() string
}

type ave struct{}
type pez struct{}
type gato struct{}

func (*ave) mover() string {
	return "Soy un ave."
}

func (*pez) mover() string {
	return "Nado , Nado, Nado"
}

func (*gato) mover() string {
	return "Caminoo"
}

func animalMove(animal Animal) {
	fmt.Println(animal.mover())
}

func main() {
	tiburoncin := pez{}
	animalMove(&tiburoncin)
}
