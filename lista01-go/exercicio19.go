package main

import (
	"fmt"
	"math"
)

func main() {
	var numero, i int
	var soma float64 = 0

	fmt.Println("Escreva um número inteiro positivo e maior que 1.")
	fmt.Scan(&numero)

	if numero > 1 {
		for i = 0; i <= numero; i++ {
			soma += 1.0 / float64(i+1)
		}
		fmt.Printf("%.6f\n", math.Round(soma*1e6)/1e6)
	} else {
		fmt.Println("Número inválido!")
	}
}
