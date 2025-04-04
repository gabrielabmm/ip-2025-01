package main

import (
	"fmt"
	"math"
)

func main() {
	var raio, altura, areaCirculo, areaLateral, areaTotal, custo float64
	const pi = 3.14159

	fmt.Print("Qual o raio da lata?\n")
	fmt.Scan(&raio)

	fmt.Print("Qual é a altura da lata?\n")
	fmt.Scan(&altura)

	areaCirculo = pi * raio * raio
	areaLateral = 2 * pi * raio * altura
	areaTotal = 2 * areaCirculo + areaLateral

	custo = 100 * areaTotal

	fmt.Printf("O VALOR DO CUSTO É= %.2f\n", math.Round(custo*100)/100)
}
