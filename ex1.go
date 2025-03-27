package main

import "fmt"

func main() {
	var matricula, i int
	var notaPr [8]float32
	var notaEx [5]float32
	var notaF, presenca, somaPr, somaEx, mediaN float32

	{
		{
			fmt.Println("Qual é o número de matrícula do aluno?")
			fmt.Scan(&matricula)

			fmt.Println("Quais são as oito notas das provas?")
			{
				for i=0; i < 8; i++ {
					fmt.Scan(&notaPr[i])
					somaPr += notaPr[i]
				}

				fmt.Println("Quais são as cinco notas obtidas pela lista de exercícios?")
				{
					for i=0; i < 5; i++{
						fmt.Scan(&notaEx[i])
						somaEx += notaEx[i]
					}

					fmt.Println("Qual a nota do trabalho final?")
					fmt.Scan(&notaF)

					fmt.Println("Qual é a quantidade de horas em que o aluno esteve presente nas aulas?")
					fmt.Scan(&presenca)

					mediaN = (0.7*somaPr/8 + 0.15*somaEx + 0.15*notaF)

					if mediaN >= 6 && presenca >= 96 {
						fmt.Println("Matrícula: ", matricula, ", Nota final:", mediaN, "Situação final: APROVADO")

					} else if mediaN >= 6 && presenca < 96 {
						fmt.Println("Matrícula: ", matricula, ", Nota final:", mediaN, "Situação final: REPROVADO POR FREQUÊNCIA")

					} else if mediaN < 6 && presenca >= 96 {
						fmt.Println("Matrícula: ", matricula, ", Nota final:", mediaN, "Situação final: REPROVADO POR NOTA")

					} else {
						fmt.Println("Matrícula: ", matricula, ", Nota final:", mediaN, "Situação final: REPROVADO POR NOTA E FREQUÊNCIA")
					}

				}