package main

import (
	"fmt"
)

func main() {
	var contaAgua, consumo, residencial, comercial, industrial float64
	var tipo string

	fmt.Print("Qual o valor da conta de água? \n")
	fmt.Scan(&contaAgua)

	fmt.Print("Qual é o consumo de água por metros cúbicos? \n")
	fmt.Scan(&consumo)

	fmt.Print("Qual o tipo de consumidor? (C - Comercial, R - Residencial, I - Industrial) \n")
	fmt.Scan(&tipo)

	fmt.Printf("CONTA= %.2f\n", contaAgua)

	residencial = 5 + 0.05*contaAgua
	comercial = 500 + 0.25*(contaAgua-80)
	industrial = 800 + 0.04*(contaAgua-100)

	switch tipo {
	case "C", "c":
		fmt.Printf("VALOR DA CONTA= %.2f\n", comercial)
	case "R", "r":
		fmt.Printf("VALOR DA CONTA= %.2f\n", residencial)
	case "I", "i":
		fmt.Printf("VALOR DA CONTA= %.2f\n", industrial)
	default:
		fmt.Println("Tipo de consumidor inválido.")
	}
}
