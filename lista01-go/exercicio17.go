package main

import (
	"fmt"
)

func main() {
	var numeroX, numeroY, i, soma int

	fmt.Println("Quais os números x e y?")
	fmt.Scan(&numeroX)
	fmt.Println()
	fmt.Scan(&numeroY)
	fmt.Println()

	soma = numeroX - 2

	for i = 0; ; i++ {
		if numeroX%2 == 0 {
			soma = soma + 2
			fmt.Printf("%d\t", soma)
			if soma >= numeroY {
				break
			}
		} else {
			fmt.Println("O PRIMEIRO NÚMERO NÃO É PAR")
			break
		}
	}
}
