package main

import "fmt"

func main(){
var X, i int
var numeros[] int
var pares[] int
var impares[] int

for i=0; i<10; i++{

fmt.Scan(&X)

numeros = append(numeros, X)

}

soma := 0
for i=0; i< 10; i++{
if numeros[i] % 2 == 0 {
pares = append(pares, numeros[i])

soma += numeros[i]
}
}
soma2 := 0
for i=0; i< 10; i++{
	if numeros[i] % 2 != 0 {
	impares = append(impares, numeros[i])
	
	soma2 += 1
	}
	}

	for i=0; i<len(pares); i++{
		fmt.Print(pares[i], " ")
	}

	fmt.Println("\n", soma)

	for i=0; i<len(impares); i++{
fmt.Print(impares[i], " ")
	}

	fmt.Println("\n", soma2)

}
