

programa
{
	inclua biblioteca Matematica --> mat
	
	funcao inicio()
	{
	
		real nota1, nota2, nota3, media

		
		escreva("Digite a primeira nota: ")
		leia(nota1)

		escreva("Digite a segunda nota: ")
		leia(nota2)

		escreva("Digite a terceira nota: ")
		leia(nota3)

		/* Calcula a média final do usuário */
		media = (nota1 + nota2 + nota3) / 3
		
		escreva ("Sua média é:", media) 
		escreva ("\n")
		se (media >= 7)
		{
		 	escreva("APROVADO")
		}
		senao
		{
			escreva("REPROVADO")
		}

		escreva("\n")
	}
}
