package main

import "fmt"

func main() {
	var B, E int

	fmt.Scan(&B, &E)

	resultado := B
	if E == 1 {
		fmt.Println(B)
	} else if E == 0 {
		fmt.Println(1)
	} else {
		for i := 1; i < E; i++ {

			resultado = resultado * B

		}
		fmt.Println(resultado)

	}
}
