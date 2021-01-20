package concesionaria

import (
	"fmt"
	"sort"
)

// OrdenarPorPrecio ...
func OrdenarPorPrecio() {
	// Guardo en los slice que voy a ordenar
	precioSorted := make([]float64, 0)
	fmt.Println("Vehiculos ordenados por precio de mayor a menor: ")
	for _, auto := range Autos {
		precioSorted = append(precioSorted, auto.Precio)
	}
	for _, moto := range Motos {
		precioSorted = append(precioSorted, moto.Precio)
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(precioSorted)))

	for i := 0; i < len(precioSorted); i++ {
		for _, auto := range Autos {
			if auto.Precio == precioSorted[i] {
				fmt.Println(auto.Marca, auto.Modelo)
			}
		}
		for _, moto := range Motos {
			if moto.Precio == precioSorted[i] {
				fmt.Println(moto.Marca, moto.Modelo)
			}
		}
	}

}
