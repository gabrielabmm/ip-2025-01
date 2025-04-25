package main

import "fmt"

func main() {
var X[]int
var i, numero int
var result string
	for i=0; i<10; i++{
fmt.Scan(&numero)

X = append(X, numero)

	}
result = "Não existe nenhum número nessa condição"
	for i=0; i< len(X); i++{
if(X[i] > 50){
	fmt.Println(X[i], i)
	result = " "
}

	}
	fmt.Println(result)
}
