package main

import (
	"fmt"
	"math"
)

func main() {
	var elementoA, elementoB, elementoC, elementoD, det float64

	fmt.Println("Quais os quatro elementos (a,b,c e d) da matriz?")
	fmt.Scan(&elementoA, &elementoB, &elementoC, &elementoD)

	det = (elementoA * elementoD) - (elementoB * elementoC)

	fmt.Printf("O VALOR DO DETERMINANTE É = %.2f\n", det)
}
