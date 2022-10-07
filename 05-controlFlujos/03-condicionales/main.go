package main

import "fmt"

func main() {
	var consumo, descuento, total, subtotal, igv float64
	var datosDescuento string

	fmt.Println("Ingrese el consumo")
	fmt.Scanln(&consumo)
	if consumo > 0 {
		if consumo <= 100 {
			descuento = consumo * 0.10
			datosDescuento = "Descuento del 10%"
		} else if consumo > 100 && consumo <= 200 {
			descuento = consumo * 0.20
			datosDescuento = "Descuento del 20%"
		} else if consumo > 200 && consumo <= 300 {
			descuento = consumo * 0.30
			datosDescuento = "Descuento del 30%"
		}
	} else {
		fmt.Println("Error al ingresar el consumo")
	}

	subtotal = consumo - descuento
	igv = subtotal * 0.19
	total = subtotal + igv
	fmt.Println("-------------Factura de consumo-------------")
	fmt.Println("Descuento que aplica: ", datosDescuento)
	fmt.Println(" valor consumo: ", consumo)
	fmt.Println(" Descuento ", descuento)
	fmt.Println(" Valor del iva aplicado", igv)
	fmt.Println(" Valor con descuento ", subtotal)

	fmt.Println("Total a pagar ", total)

}
