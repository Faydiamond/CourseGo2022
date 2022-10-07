package figuras

import "fmt"

type Calculos interface {
	perimetro() float32
	area() float32
}

func CalcularValor(calculo Calculos) {
	fmt.Println("El perimetro es: ", calculo.perimetro())
	fmt.Println("El area es: ", calculo.area())
}
