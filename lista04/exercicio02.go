package main

import "fmt"

func main() {
	var X []int
	var Y []int
	var soma01 []int
	var soma02 []int
	var i, numero, numero2 int

	for i = 0; i < 10; i++ {
		fmt.Scan(&numero)

		X = append(X, numero)
	}
	for i = 0; i < 5; i++ {
		fmt.Scan(&numero2)

		Y = append(Y, numero2)
	}

	soma := 0
	soma2 := 0

	for i = 0; i < len(Y); i++ {
		soma += Y[i]
		soma2 += Y[i]
	}

	for i = 0; i < len(X); i++ {
			if X[i]%2 == 0 {soma01 = append(soma01, (soma + X[i]))
		}else {soma02 = append(soma02, (soma2 + X[i]))}

	}

	for i = 0; i < len(soma01); i++ {
		fmt.Println(soma01[i])
	}

	for i = 0; i < len(soma02); i++ {
		fmt.Println(soma02[i])
	}

}
