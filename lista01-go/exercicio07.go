package main

import (
	"fmt"
	"math"
)

func main() {
	var valorF, valorC, conversaoFC, conversaoPM float64

	fmt.Print("Qual o valor, em Fahrenheit, que será convertido para Celsius? \n")
	fmt.Scan(&valorF)

	fmt.Print("Qual é a quantidade de chuva, em polegadas, que será convertida para milímetros? \n")
	fmt.Scan(&valorC)

	conversaoFC = (5*valorF - 160) / 9

	conversaoPM = valorC * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", math.Round(conversaoFC*100)/100)
	fmt.Printf("A QUANTIDADE DE CHUVA É = %.2f\n", math.Round(conversaoPM*100)/100)
}
