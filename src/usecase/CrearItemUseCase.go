package usecase

import "workshop/src/domain"

type CrearItemUseCase struct {
	repositorio domain.Repositorio
}

func (useCase *CrearItemUseCase) SetRepositorio(repositorio domain.Repositorio) {
	useCase.repositorio = repositorio
}

func (useCase *CrearItemUseCase) Ejecutar(nombre string, precio float32) {

	item := domain.Item{
		Nombre: nombre,
		Precio: precio,
	}

	useCase.repositorio.CrearItem(item)
}
