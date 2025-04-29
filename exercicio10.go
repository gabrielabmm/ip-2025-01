package main

import (
	"fmt"
	"sort"
)

func main() {
var P_inicial, D, DeltaP, P_min, i, lucro, Q_inicial, DeltaQ, precox, lucrox, ingressosx float64
var preco[] float64
var qntd[] float64
var lucroobt[] float64
var lucroobt2[] float64
var tamanho int

fmt.Scan(&P_inicial, &Q_inicial, &D, &DeltaP, &DeltaQ, &P_min)

for i=0; P_inicial - DeltaP*i >= P_min; i++{

preco = append(preco, (P_inicial - DeltaP * i))
qntd = append(qntd, (Q_inicial + DeltaQ * i))

	lucro = ((P_inicial - DeltaP * i) * (Q_inicial + DeltaQ * i)) - D

lucroobt = append(lucroobt, lucro)
lucroobt2 = append(lucroobt2, lucro)

tamanho ++

}

sort.Float64s(lucroobt2)


for j:=0; j< tamanho; j++{

if (lucroobt[j] == lucroobt2[tamanho-1]){
precox = P_inicial - DeltaP * float64(j)
lucrox = ((P_inicial - DeltaP * float64(j)) * (Q_inicial + DeltaQ * float64(j))) - D
ingressosx = Q_inicial + DeltaQ * float64(j)
break
}
}


fmt.Printf("%-20s %-20s %-20s\n", "Preco", "Ingressos Vendidos", "Lucro")

for q := 0; q < tamanho; q++ {
	fmt.Printf("%-20.2f %-20.f %-20.2f\n", preco[q], qntd[q], lucroobt[q])
}

fmt.Printf("Lucro maximo: %.2f\n", lucrox)
fmt.Printf("Na faixa de preco: %.2f com %.f ingressos.\n", precox, ingressosx)
			
}
