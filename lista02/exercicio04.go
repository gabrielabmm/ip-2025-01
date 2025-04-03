package main

import "fmt"

func main() {
	var numN, valores, valorIn, incremento, i float64
	var soma, subtracao, divisao, multiplicacao float64
	var somas, subtracoes, multiplicacoes, divisoes []float64

	fmt.Println("Escreva um número de 0 a 9.")
	fmt.Scan(&numN)

	fmt.Println("O gerador iniciará em qual valor?")
	fmt.Scan(&valorIn)

	fmt.Println("Digite a quantidade de valores.")
	fmt.Scan(&valores)

	fmt.Println("Qual será o valor do incremento?")
	fmt.Scan(&incremento)

	i = 1

	for i <= valores {
		soma = numN + (valorIn + incremento*i)

		subtracao = numN - (valorIn + incremento*i)

		multiplicacao = numN * (valorIn + incremento*i)

		divisao = numN / (valorIn + incremento*i)

		i++
	}

	{

		somas = append(somas, soma)
		subtracoes = append(subtracoes, subtracao)
		multiplicacoes = append(multiplicacoes, multiplicacao)
		divisoes = append(divisoes, divisao)
	}

	for idx := range somas {
		fmt.Println("Tabuada de soma:", somas[idx])
	}

	for idx := range subtracoes {
		fmt.Println("Tabuada de subtração:", subtracoes[idx])
	}

	for idx := range multiplicacoes {
		fmt.Println("Tabuada de multiplicação:", multiplicacoes[idx])
	}

	for idx := range divisoes {
		fmt.Println("Tabuada de divisão:", divisoes[idx])
	}

}
