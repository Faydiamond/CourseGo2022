package main

import "fmt"

type heroe struct {
	Name string
	Type string
	Age  int
}

func main() {
	thor := heroe{Name: "Thor", Age: 29, Type: "Electrico"}
	fmt.Println(thor)

}
