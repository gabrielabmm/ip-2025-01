package main

import "fmt"

func main() {
	var i, N int
	var n1, n2, media float64
	var notas []float64

	fmt.Scan(&N)

	for i = 1; i <= N; i++ {
		fmt.Scan(&n1, &n2)

		media = (n1 + n2) / 2

		notas = append(notas, media)

	}
	aprovados := 0
	exame := 0
	reprovados := 0

	for i, valor := range notas {
		aluno := i + 1
		if valor > 7 {
			fmt.Printf("Aluno %d: Aprovado\n", aluno)
			aprovados += 1

		} else if valor >= 3 && valor <= 7 {
			fmt.Println("Aluno %d: Exame", aluno)
			exame += 1

		} else {
			fmt.Println("Aluno %d: Reprovado", aluno)
			reprovados += 1

		}
	}
	fmt.Printf("Total Aprovados: %d\n", aprovados)
	fmt.Printf("Total Exame: %d\n", exame)
	fmt.Println("Total Reprovados: %d\n", reprovados)

}
