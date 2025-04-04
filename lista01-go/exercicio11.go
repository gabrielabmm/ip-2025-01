package main

import (
	"fmt"
)

func main() {
	var numInt int

	fmt.Println("Escreva um número inteiro")
	fmt.Scan(&numInt)

	if numInt%3 == 0 && numInt%5 == 0 {
		fmt.Println("O NÚMERO É DIVISÍVEL")
	} else {
		fmt.Println("O NÚMERO NÃO É DIVISÍVEL")
	}
}
