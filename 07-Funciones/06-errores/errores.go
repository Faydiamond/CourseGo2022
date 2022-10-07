package main

import (
	"errors"
	"fmt"
)

func splitt(dividiendo, divisor int) (int, error) {

	if divisor == 0 {
		return 0, errors.New("No es posible dividir entre cero")
	} else {
		return dividiendo / divisor, nil
	}

}

func main() {
	result, error := splitt(4, 1)

	if error == nil {
		fmt.Println("El resultado de la division es: ", result)
	} else {
		fmt.Println("El error es : ", error)
	}

}
