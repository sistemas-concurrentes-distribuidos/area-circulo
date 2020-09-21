package main

import (
	"fmt"
	"math"
)

func main() {
	var radio float64

	fmt.Printf("AREA DEL CIRCULO \n\n")
	fmt.Printf("Ingresa el RADIO del circulo: ")
	fmt.Scanf("%f", &radio)

	area := math.Pi * (math.Pow(radio, 2))

	fmt.Println("Area: ", area)
}
