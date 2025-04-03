package main

import (
	"fmt"
)

func main() {

	var nota1, nota2, nota3 float64

	fmt.Println("Digite a primeira nota: ")
	fmt.Scan(&nota1)

	fmt.Println("Digite a segunda nota: ")
	fmt.Scan(&nota2)

	fmt.Println("Digite a terceira nota: ")
	fmt.Scan(&nota3)

	media := (nota1 + nota2 + nota3) / 3

	fmt.Printf("Sua média é: %0.2f ", media)
	fmt.Printf("\n")

	if media >= 6 {
		fmt.Printf(" APROVADO")

	} else {
		fmt.Printf(" REPROVADO")
	}

}
