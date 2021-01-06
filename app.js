const concesionaria = [{
        "Marca": "Peugeot",
        "Modelo": 206,
        "Puertas": 4,
        "Precio": 200000
    },
    {
        "Marca": "Honda",
        "Modelo": "Titan",
        "Cilindrada": "125cc",
        "Precio": 60000
    },
    {
        "Marca": "Peugeot",
        "Modelo": 208,
        "Puertas": 5,
        "Precio": 250000
    },
    {
        "Marca": "Yamaha",
        "Modelo": "YBR",
        "Cilindrada": "160cc",
        "Precio": 80500.50
    }
]

function mostrarVehiculos() {
    concesionaria.forEach(product => {
        if (product.Puertas === undefined) {
            console.log(`Marca: ${product.Marca} // Modelo: ${product.Modelo} // Cilindrada: ${product.Cilindrada} // Precio: $${product.Precio}`)
        } else {
            console.log(`Marca: ${product.Marca} // Modelo: ${product.Modelo} // Puertas: ${product.Puertas} // Precio: $${product.Precio}`)
        }

    })
}

function filtroVehiculos() {
    let masCaro = 0
    let nombreMas = ""
    let masBarato = 500000
    let nombreMenos = ""
    let ybr = ""
    concesionaria.forEach(vehiculo => {
        if (vehiculo.Precio > masCaro) {
            masCaro = vehiculo.Precio
            nombreMas = `Vehiculo más caro: ${vehiculo.Marca} ${vehiculo.Modelo}`
        }
        if (vehiculo.Precio < masBarato) {
            masBarato = vehiculo.Precio
            nombreMenos = `Vehiculo mas barato: ${vehiculo.Marca} ${vehiculo.Modelo}`
        }
        for (let i = 0; i <= vehiculo.Modelo.length; i++) {
            if (vehiculo.Modelo.charAt(i) == 'Y') {
                ybr = `Vehiculo que contiene en el modelo la letra 'Y': ${vehiculo.Marca} ${vehiculo.Modelo} $${vehiculo.Precio}`
            }
        }
    })
    console.log(nombreMas)
    console.log(nombreMenos)
    console.log(ybr)

}

function ordenar() {
    concesionaria.sort((a, b) => b.Precio - a.Precio)
    console.log("Vehiculos ordenados por precio de mayor a menor:")
    concesionaria.forEach(vehiculo => {
        console.log(`${vehiculo.Marca} ${vehiculo.Modelo}`)
    })
}

mostrarVehiculos()
console.log("=================================")
filtroVehiculos()
console.log("=================================")
ordenar()