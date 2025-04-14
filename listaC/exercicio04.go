package main

import "fmt"

func main() {
	var numeros []int
	var numero int

	for {
		fmt.Scan(&numero)

		if numero == 0 {
			break
		}
		numeros = append(numeros, numero)
	}

	for _, valor := range numeros {
		if QuadradoPerf(valor) {
			fmt.Println(valor, "eh quadrado perfeito")
		} else {
			fmt.Println(valor, "nao eh quadrado perfeito")
		}
	}
}

func QuadradoPerf(n int) bool {
	for i := 1; i*i <= n; i++ {
		quadrado := i * i
		if quadrado == n {
			return true
		}
	}
	return false

}
