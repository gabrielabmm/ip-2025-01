package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&x)

	fmt.Printf("%d pratos de trigo pra %d tigres tristes \n", x, x)
}
