package main

import "fmt"

type user struct {
	name  string
	age   int
	pass  string
	userr string
}

type student struct {
	User user
	cod  string
}

func main() {
	// relacion uno a uno
	alex := user{
		name:  "Alex",
		age:   25,
		pass:  "dsf",
		userr: "Alex",
	}

	mariano := user{
		name:  "Mariano",
		age:   25,
		pass:  "dsf",
		userr: "Mariano",
	}

	fmt.Println(alex)
	fmt.Println(mariano)

	marianoStudent := student{
		User: alex,
		cod:  "544545455452",
	}

	fmt.Println(marianoStudent)
	fmt.Println(marianoStudent.User.name)
}
