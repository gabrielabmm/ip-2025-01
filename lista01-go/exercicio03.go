package main

import (
	"fmt"
	"math"
)

func main() {
	var n1, n2, n3, potencia, numComposicao float64

	fmt.Println("Qual é o primeiro número?")
	fmt.Scan(&n1)

	fmt.Println("Qual é o segundo número?")
	fmt.Scan(&n2)

	fmt.Println("Qual é o terceiro número?")
	fmt.Scan(&n3)

	if n1 >= 0 && n1 < 10 && n2 >= 0 && n2 < 10 && n3 >= 0 && n3 < 10 {
		numComposicao = (n1*100 + n2*10 + n3)
		potencia = math.Pow(numComposicao, 2)

		fmt.Println("O número resultado da composição é:", numComposicao)
		fmt.Printf("\n")
		fmt.Println("O resultado do número elevado ao quadrado é:", potencia)

	} else {
		fmt.Println("DIGITO INVALIDO")
	}

}
