package main

import (
	"fmt"
	"math"
)

func main() {
	var numero, numeroQuad, i, ordemN, numeroB, soma int

	fmt.Println("Qual é o número inteiro N?")
	fmt.Scan(&numero)

	numeroB = numero / 2

	for i = 1; i <= numeroB; i++ {
		soma = soma + 1
		ordemN = soma * 2
		numeroQuad = int(math.Pow(float64(ordemN), 2.0))
		fmt.Printf("%d^2 = %d\n", ordemN, numeroQuad)
	}
}
