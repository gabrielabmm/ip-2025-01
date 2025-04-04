package main

import (
	"fmt"
	"math"
)

func main() {
	var qntdTemp, i, t, conversao float64

	fmt.Print("Quantas temperaturas irão ser convertidas de Fahrenheit para Celsius? \n")
	fmt.Scan(&qntdTemp)


	for i = 1; i <= qntdTemp; i++ {

		fmt.Printf("Qual é a temperatura %.0f?\n", i)
		fmt.Scan(&t)

		conversao = (5 * (t - 32)) / 9

		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", t, math.Round(conversao*100)/100)
	}
}
