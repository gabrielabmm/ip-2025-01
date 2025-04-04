package main

import (
	"fmt"
)

func main() {
	var horas, valor, preco int

	fmt.Println("Por quantas horas a charrete foi utilizada?")
	fmt.Scan(&horas)

	preco = 5 * horas

	if horas%3 == 0 {
		valor = preco - 5
	} else {
		valor = preco
	}

	fmt.Printf("O VALOR A PAGAR É = %d\n", valor)
}

