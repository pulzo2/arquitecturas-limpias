package domain

import "fmt"

type NuevoMedio struct {
	Nombre string
}

func (nm *NuevoMedio) Pagar(monto float32, items []Item) {
	fmt.Println("****** MEDIO DE PAGO NUEVO ******")

	for _, item := range items {
		fmt.Println(item.Nombre, " - ", item.Precio)
	}

	fmt.Println("**********************************")
	fmt.Println("Total: ", monto)
}
