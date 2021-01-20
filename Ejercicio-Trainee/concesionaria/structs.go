package concesionaria

// Auto ...
type Auto struct {
	Marca   string
	Modelo  int
	Puertas int
	Precio  float64
}

// Moto ...
type Moto struct {
	Marca      string
	Modelo     string
	Cilindrada string
	Precio     float64
}

// Autos es un  slice
var Autos []Auto

// Motos es un slice
var Motos []Moto

// AddAuto agrega el auto al slice
func AddAuto(auto Auto) {
	Autos = append(Autos, auto)
}

// AddMoto agrega la moto al slice
func AddMoto(moto Moto) {
	Motos = append(Motos, moto)
}
