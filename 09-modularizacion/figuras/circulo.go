package figuras

type Circulo struct {
	Radio float32
}

func (c *Circulo) perimetro() float32 {
	return 2 * 3.1415 * c.Radio
}

func (c *Circulo) area() float32 {
	return 3.1415 * (c.Radio * c.Radio)
}
