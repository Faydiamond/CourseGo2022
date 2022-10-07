package main

import "fmt"

func main() {
	computer := make(map[string]float32)
	fmt.Println(computer)

	computer["monitor"] = 850.589
	computer["mouse"] = 50.589
	computer["teclado"] = 70.789
	computer["ram"] = 285.789
	fmt.Println(computer)
	fmt.Println("El precio de la ram es de: ", computer["ram"])

	plants := make(map[string][]string)
	fmt.Println(plants)
	plants["buganvilla"] = []string{"Morada", "Rosada", "Blanca", "Naranja"}
	plants["hortensias"] = []string{"Azules", "Moradas", "Blancas"}
	plants["geranias"] = []string{"Vino tinto", "Rosado", "Morado"}
	fmt.Println(plants)

	chamarra := make(map[string][]string)
	chamarra["material"] = []string{"Cuero", "150.000"}
	chamarra["herrajes"] = []string{"acero", "15.000"}
	chamarra["forro"] = []string{"algodon", "10.000"}

	fmt.Println(chamarra)
}
