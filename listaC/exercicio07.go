package main

import (
	"fmt"
	"math"
)

func main() {
	var X, i, somatorio, fatorial float64

	fmt.Scan(&X)

	fatorial = 1
	somatorio = X
	for i = 1; i <= 20; i++ {
		fatorial *= i

		somatorio += math.Pow((-X), i) / fatorial

	}
	fmt.Println(somatorio)
}
