package main

import "fmt"

type Calculos interface {
	perimetro() float32
	area() float32
}

type Cuadrado struct {
	ancho  float32
	altura float32
}
type Circulo struct {
	radio float32
}

func (c *Cuadrado) perimetro() float32 {
	return 2*c.ancho + 2*c.altura
}

func (c *Cuadrado) area() float32 {
	return c.ancho * c.altura
}

func (c *Circulo) perimetro() float32 {
	return 2 * 3.1415 * c.radio
}

func (c *Circulo) area() float32 {
	return 3.1415 * (c.radio * c.radio)
}

func calcularValor(calculo Calculos) {
	fmt.Println("El perimetro es: ", calculo.perimetro())
	fmt.Println("El area es: ", calculo.area())
}

func main() {
	cuadro01 := Cuadrado{5, 5}
	calcularValor(&cuadro01)

	circulo01 := Circulo{15}
	calcularValor(&circulo01)
}
