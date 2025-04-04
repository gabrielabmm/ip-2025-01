package main

import (
	"fmt"
)

func main() {
	var valorIn, razao, numElem, soma float64
	var i int

	fmt.Println("Qual o valor inicial da progressão aritmética?")
	fmt.Scan(&valorIn)

	fmt.Println("Qual a razão da progressão?")
	fmt.Scan(&razao)

	fmt.Println("Quantos elementos formam a progressão?")
	fmt.Scan(&numElem)

	soma = 0
	for i = 0; float64(i) < numElem; i++ {
		soma += valorIn + razao*float64(i)
	}

	fmt.Printf("O valor da soma dos %.0f elementos é = %.2f\n", numElem, soma)
}
