package main

import "fmt"

type phone struct {
	mark  string
	model string
}

func (p *phone) showdata() {
	fmt.Println(p.mark, "  -  ", p.model)
}

func main() {
	fmt.Println("Hola")
	phone1 := phone{"Xiaomi", "Note 8 Pro"}
	phone1.showdata()
	phone1.mark = "Apple"
	phone1.model = "se"
	phone1.showdata()

	phone2 := phone{mark: "Lg", model: "Optimus"}
	phone2.showdata()
}
