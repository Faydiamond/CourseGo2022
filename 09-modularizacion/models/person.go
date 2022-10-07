package models

type Person struct {
	name string
	age  int
}

func (p *Person) Construct(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) GetNombre() string {
	return p.name
}

func (p *Person) SetNombre(name string) {
	p.name = name
}

func (p *Person) GetEdad(age int) int {
	return p.age
}

func (p *Person) SetEdad(age int) {
	p.age = age
}
