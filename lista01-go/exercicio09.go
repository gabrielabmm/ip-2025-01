package main

import (
	"fmt"
	"math"
)

func main() {
	var coefA, coefB, coefC, delta float64

	fmt.Println("Quais os coeficientes A, B e C?")
	fmt.Scan(&coefA, &coefB, &coefC)

	delta = math.Pow(coefB, 2) - 4*coefA*coefC

	fmt.Printf("O VALOR DE DELTA É = %.2f\n", delta)
}
