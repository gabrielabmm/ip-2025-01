package main

import (
	"fmt"
	"math"
)

func main() {
	var altura, aresta, volume, areaHex float64

	fmt.Println("Qual é a altura da pirâmide?")
	fmt.Scan(&altura)

	fmt.Println("\nQual número corresponde à aresta do hexágono que forma a base da pirâmide?")
	fmt.Scan(&aresta)

	areaHex = (3 * math.Pow(aresta, 2) * math.Sqrt(3)) / 2
	volume = (areaHex * altura) / 3

	fmt.Printf("O VOLUME DA PIRÂMIDE = %.2f METROS CUBICOS\n", volume)
}
