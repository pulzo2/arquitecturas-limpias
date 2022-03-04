package domain

type Repositorio interface {
	CrearItem(item Item)
	ListarItems() []Item
}
