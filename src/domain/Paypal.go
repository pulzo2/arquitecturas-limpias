package domain

import "fmt"

type Paypal struct {
	Email    string
	Password string
}

func (p *Paypal) Pagar(monto float32, items []Item) {

	fmt.Println("****** MEDIO DE PAGO PAYPAL ******")

	for _, item := range items {
		fmt.Println(item.Nombre, " - ", item.Precio)
	}

	fmt.Println("**********************************")
	fmt.Println("Total: ", monto)
}
