package domain

type OrdenCompra struct {
	items []Item
}

func (oc *OrdenCompra) AgregarItem(item Item) {
	oc.items = append(oc.items, item)
}

func (oc *OrdenCompra) CalcularTotal() float32 {
	var total float32 = 0
	for _, item := range oc.items {
		total = total + item.Precio
	}
	return total
}

func (oc *OrdenCompra) Pagar(medioDePago MedioDePago) {
	monto := oc.CalcularTotal()
	medioDePago.Pagar(monto, oc.items)
}
