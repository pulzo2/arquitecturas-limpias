package domain

type MedioDePago interface {
	Pagar(monto float32, items []Item)
}
