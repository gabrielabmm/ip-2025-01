programa {
  funcao inicio() {
real nota

    escreva ("Qual é a nota? \n")
    leia (nota)

    se (nota >= 9 e nota <= 10)
    {escreva ("NOTA = ", nota, "\n", "CONCEITO = A")}
   
    senao se (nota >=7.5 e nota < 9)
    {escreva ("NOTA = ", nota, "\n", "CONCEITO = B")}
   
    senao se (nota >= 6 e nota < 7.5)
    {escreva ("NOTA = ", nota, "\n", "CONCEITO = C")}
 
     senao se (nota >= 0 e nota < 6)
    {escreva ("NOTA = ", nota, "\n", "CONCEITO = D")}

    senao {escreva ("NOTA INVÁLIDA")}

  }

  
}
