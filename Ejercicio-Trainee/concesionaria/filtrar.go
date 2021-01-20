package concesionaria

import (
	"fmt"
	"strconv"
	"strings"
)

// FiltrarPorPrecios ..
func FiltrarPorPrecios() {
	mas := 0.0
	nombreMas := ""
	modeloMas := ""
	menos := 9999999.9
	nombreMenos := ""
	modeloMenos := ""
	for _, auto := range Autos {
		for _, moto := range Motos {
			if (auto.Precio > mas) && (moto.Precio > mas) {
				if auto.Precio > moto.Precio {
					mas = auto.Precio
					nombreMas = auto.Marca
					modeloMas = strconv.Itoa(auto.Modelo)
				} else {
					mas = moto.Precio
					nombreMas = moto.Marca
					modeloMas = moto.Modelo
				}
			}
			if (auto.Precio < menos) && (moto.Precio < menos) {
				if auto.Precio < moto.Precio {
					menos = auto.Precio
					nombreMenos = auto.Marca
					modeloMenos = strconv.Itoa(auto.Modelo)
				} else {
					menos = moto.Precio
					nombreMenos = moto.Marca
					modeloMenos = moto.Modelo
				}
			}
		}
	}
	fmt.Println("Vehiculo más caro: ", nombreMas, modeloMas)
	fmt.Println("Vehiculo más barato: ", nombreMenos, modeloMenos)
}

// FiltarLetra ...
func FiltarLetra() {
	encontrado := false
	marca := ""
	modelo := ""
	precio := 0.0
	for _, auto := range Autos {
		for _, moto := range Motos {
			s := strings.Split(strconv.Itoa(auto.Modelo), "")
			s1 := strings.Split(moto.Modelo, "")
			for _, letra := range s {
				for _, letr := range s1 {
					if letra == "Y" {
						fmt.Println("Encontrado")
						encontrado = true
						marca = auto.Marca
						modelo = strconv.Itoa(auto.Modelo)
						precio = auto.Precio
					}
					if letr == "Y" {
						// fmt.Println("Encontrado en motos")
						encontrado = true
						marca = moto.Marca
						modelo = moto.Modelo
						precio = moto.Precio
					}
				}
			}
		}
	}
	if encontrado {
		fmt.Printf("Vehiculo que contiene en el modelo la letra 'Y': %v %v $%v\n", marca, modelo, precio)
	}
}
