package main

import (
	"fmt"
	"workshop/src/infraestructure/repositorio"
	"workshop/src/usecase"
)

func main() {

	repositorio := &repositorio.Memoria{}
	crearItemUseCase := usecase.CrearItemUseCase{}
	crearItemUseCase.SetRepositorio(repositorio)

	crearItemUseCase.Ejecutar("Item 1", 100)
	crearItemUseCase.Ejecutar("Item 2", 50)
	crearItemUseCase.Ejecutar("Item 3", 25)
	crearItemUseCase.Ejecutar("Item 4", 78)
	crearItemUseCase.Ejecutar("Item 5", 108)

	listarItems := usecase.ListarItemsUseCase{}
	listarItems.SetRepositorio(repositorio)
	items := listarItems.Ejecutar()
	fmt.Println(items)

	// item1 := domain.Item{Nombre: "Item 1", Precio: 100}
	// item2 := domain.Item{Nombre: "Item 2", Precio: 55.5}
	// item3 := domain.Item{Nombre: "Item 3", Precio: 19}

	// ordenCompra := domain.OrdenCompra{}
	// ordenCompra.AgregarItem(item1)
	// ordenCompra.AgregarItem(item2)
	// ordenCompra.AgregarItem(item3)

	// payPal := domain.Paypal{
	// 	Email:    "enoc.martinez@pulzo",
	// 	Password: "123",
	// }

	// tarjetaCredito := domain.TarjetaCredito{
	// 	Nombre: "Enoc",
	// 	Numero: "1092-2123-1234-2343",
	// 	Ccv:    "123",
	// 	FecExp: "04/29",
	// }

	// nuevoMedio := domain.NuevoMedio{
	// 	Nombre: "Enoc",
	// }

	// ordenCompra.Pagar(&nuevoMedio)
}
