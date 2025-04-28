package main

import (
	"fmt"
	"sort"
)

func main() {
var P_inicial, D, DeltaP, P_min, i, lucro, Q_inicial, DeltaQ float64
var preco[] float64
var qntd[] float64
var lucroobt[] float64
var lucroobt2[] float64
var tamanho int

fmt.Scan(&P_inicial, &D, &DeltaP, &P_min, &Q_inicial, &DeltaQ)

for i=0; i< i+1; i++{
if (P_inicial - DeltaP * i <= P_min){

preco = append(preco, (P_inicial - DeltaP * i))
qntd = append(qntd, (Q_inicial + DeltaQ * i))

	lucro = ((P_inicial - DeltaP * i) * (Q_inicial + DeltaQ * i)) - D

lucroobt = append(lucroobt, lucro)
lucroobt2 = append(lucroobt2, lucro)

tamanho += 1
}

}

sort.Float64s(lucroobt2)


for j:=0; j< tamanho; j++{

if (lucroobt[j] == lucroobt2[0]){
indice := j
}
}

fmt.Printf("%-10s %-10s %-10s\n", "Preco", "Ingressos Vendidos", "Lucro")

for i=0; i< tamanho; i++{
	fmt.Printf("%-10d \n", preco[i])}

	for i=0; i< tamanho; i++{
		fmt.Printf("%-10d \n", qntd[i])}

		for i=0; i< tamanho; i++{
			fmt.Printf("%-10d \n", lucroobt[i])}
			
			fmt.Printf("Preco min %-10d" preco[indice])
}
