package main

import (
	"fmt"
	"math"
)

func main() {
	var salario, salarioReaj float64

	fmt.Println("Qual é o salário do funcionário?")
	fmt.Scan(&salario)

	if salario < 300 {
		salarioReaj = salario * 1.5
	} else {
		salarioReaj = salario * 1.3
	}

	fmt.Printf("SALÁRIO COM REAJUSTE = %.2f\n", math.Round(salarioReaj*100)/100)
}
