package main

import "fmt"

func main() {

	var salarioMin float64
	var kwGastos int

	fmt.Println("Qual é o valor do salário mínimo? ")
	fmt.Scan(&salarioMin)

	fmt.Println("Quantos kW foram gastos pela residência? ")
	fmt.Scan(&kwGastos)

	kWunico := (salarioMin * 0.7)
	valorkW := kWunico / 100       

	fmt.Printf("O valor, em reais, de cada kW é: R$ %.2f\n", valorkW)

	valorRes := valorkW * float64(kwGastos)

	fmt.Printf("O valor, em reais, a ser pago pela residência é: R$ %.2f\n", valorRes)

	novoValor := valorRes * 0.9 // 10% de desconto

	fmt.Printf("O novo valor a ser pago, com desconto de 10%%, é: R$ %.2f\n", novoValor)
}

