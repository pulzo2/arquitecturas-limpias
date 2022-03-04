package repositorio

import "workshop/src/domain"

type Memoria struct {
	items []domain.Item
}

func (m *Memoria) CrearItem(item domain.Item) {
	m.items = append(m.items, item)
}

func (m *Memoria) ListarItems() []domain.Item {
	return m.items
}
