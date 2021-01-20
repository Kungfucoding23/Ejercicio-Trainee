package concesionaria

import "fmt"

// Imprimir ...
func Imprimir() {
	for _, auto := range Autos {
		fmt.Printf("Marca: %v // Modelo: %v // Puertas: %v // Precio: $%v\n", auto.Marca, auto.Modelo, auto.Puertas, auto.Precio)
	}
	for _, moto := range Motos {
		fmt.Printf("Marca: %v // Modelo: %v // Cilindrada: %v // Precio: $%v\n", moto.Marca, moto.Modelo, moto.Cilindrada, moto.Precio)
	}
}
