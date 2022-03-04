package domain

import "fmt"

type TarjetaCredito struct {
	Nombre string
	Numero string
	Ccv    string
	FecExp string
}

func (tc *TarjetaCredito) Pagar(monto float32, items []Item) {

	fmt.Println("****** MEDIO DE TARJETA DE CREDITO ******")

	for _, item := range items {
		fmt.Println(item.Nombre, " - ", item.Precio)
	}

	fmt.Println("**********************************")
	fmt.Println("Total: ", monto)
}
