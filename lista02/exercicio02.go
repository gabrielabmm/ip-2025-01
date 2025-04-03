package main

import "fmt"

func main() {
	var popA, popB float64
	var taxaA float64 = 0.03
	var taxaB float64 = 0.015
	var i float64 = 0

	fmt.Println("Qual a população do país A?")
	fmt.Scan(&popA)

	fmt.Println("Qual a população do país B?")
	fmt.Scan(&popB)

	i = 0
	for popA > popB {
		popA += (taxaA * popA)
		popB += (taxaB * popB)
		i++

	}
	{
		fmt.Println(i)
	}

}
