package main

import (
	"fmt"
)

func main() {
	var nota float64

	fmt.Println("Qual é a nota?")
	fmt.Scan(&nota)

	if nota >= 9 && nota <= 10 {
		fmt.Printf("NOTA = %.2f\nCONECITO = A\n", nota)
	} else if nota >= 7.5 && nota < 9 {
		fmt.Printf("NOTA = %.2f\nCONECITO = B\n", nota)
	} else if nota >= 6 && nota < 7.5 {
		fmt.Printf("NOTA = %.2f\nCONECITO = C\n", nota)
	} else if nota >= 0 && nota < 6 {
		fmt.Printf("NOTA = %.2f\nCONECITO = D\n", nota)
	} else {
		fmt.Println("NOTA INVÁLIDA")
	}
}
