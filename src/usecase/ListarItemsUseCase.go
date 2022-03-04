package usecase

import "workshop/src/domain"

type ListarItemsUseCase struct {
	repositorio domain.Repositorio
}

func (useCase *ListarItemsUseCase) SetRepositorio(repositorio domain.Repositorio) {
	useCase.repositorio = repositorio
}

func (useCase *ListarItemsUseCase) Ejecutar() []domain.Item {

	return useCase.repositorio.ListarItems()
}
