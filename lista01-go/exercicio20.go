package main

import (
	"fmt"
)

func main() {
	var horas, min, seg, tempo float64

	fmt.Println("Qual o valor do tempo expresso em horas?")
	fmt.Scan(&horas)

	fmt.Println("Qual o valor do tempo expresso em minutos?")
	fmt.Scan(&min)

	fmt.Println("Qual o valor do tempo expresso em segundos?")
	fmt.Scan(&seg)

	tempo = (horas * 3600) + (min * 60) + seg

	fmt.Printf("O TEMPO EM SEGUNDOS É = %.0f\n", tempo)
}
