package main

import (
	"fmt"
)
func main() {
	var numJogos, i int
	var numPessoas, percPo, percGe, percAr, percCa, renda1, renda2, renda3, renda4, rendaTotal float64

	fmt.Println("Qual o número de casos de teste?")
	fmt.Scan(&numJogos)
	fmt.Printf("\n")

	for i = 1; i <= numJogos; i++ {

		fmt.Println("Quantas pessoas compraram ingresso para o jogo ", i, "?")
		fmt.Scan(&numPessoas)

		fmt.Println("Qual a percentagem de pessoas que compraram na categoria Popular?")
		fmt.Scan(&percPo)

		fmt.Println("Qual a percentagem de pessoas que compraram na categoria Geral?")
		fmt.Scan(&percGe)

		fmt.Println("Qual a percentagem de pessoas que compraram na categoria Arquibancada?")
		fmt.Scan(&percAr)

		fmt.Println("Qual a percentagem de pessoas que compraram na categoria Cadeiras?")
		fmt.Scan(&percCa)

		renda1 = (numPessoas * percPo) / 100
		renda2 = (numPessoas * percGe * 5) / 100
		renda3 = (numPessoas * percAr * 10) / 100
		renda4 = (numPessoas * percCa * 20) / 100

		rendaTotal = renda1 + renda2 + renda3 + renda4

		fmt.Printf("A renda do jogo é %0.2f", rendaTotal)
	}
}
