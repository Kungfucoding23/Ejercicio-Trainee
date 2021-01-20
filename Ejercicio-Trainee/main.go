package main

import (
	"fmt"
	"trainee/concesionaria"
)

func main() {

	// // Open the file
	// file, err := os.Open("concesionaria.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// // create a new decoder
	// var concesionariaDecoder *json.Decoder = json.NewDecoder(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // initialize the storage for the decoded data
	// var concesionariaList []concesionaria.Auto

	// // decode the data
	// err = concesionariaDecoder.Decode(&concesionariaList)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, vehiculo := range concesionariaList {
	// 	fmt.Println(vehiculo.Marca)
	// 	fmt.Println(vehiculo.Modelo)
	// 	fmt.Println(vehiculo.Puertas)
	// 	fmt.Println(vehiculo.Precio)

	// }

	auto := concesionaria.Auto{
		Marca:   "Peugeot",
		Modelo:  206,
		Puertas: 4,
		Precio:  200000,
	}
	moto := concesionaria.Moto{
		Marca:      "Honda",
		Modelo:     "Titan",
		Cilindrada: "123cc",
		Precio:     60000,
	}
	auto2 := concesionaria.Auto{
		Marca:   "Peugeot",
		Modelo:  208,
		Puertas: 5,
		Precio:  250000,
	}
	moto2 := concesionaria.Moto{
		Marca:      "Yamaha",
		Modelo:     "YBR",
		Cilindrada: "160cc",
		Precio:     80500.50,
	}

	concesionaria.AddAuto(auto)
	concesionaria.AddAuto(auto2)
	concesionaria.AddMoto(moto)
	concesionaria.AddMoto(moto2)

	concesionaria.Imprimir()
	fmt.Println("===========================================")
	concesionaria.FiltrarPorPrecios()
	concesionaria.FiltarLetra()
	fmt.Println("===========================================")
	concesionaria.OrdenarPorPrecio()

}
