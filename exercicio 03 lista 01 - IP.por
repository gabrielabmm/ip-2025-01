programa {
  	inclua biblioteca Matematica
  funcao inicio() {
inteiro n1, n2, n3, potencia, numComposicao

    escreva ("Qual é o primeiro número?\n")
    leia (n1)

   escreva ("Qual é o segundo número?\n")
    leia (n2)

       escreva ("Qual é o terceiro número?\n")
    leia (n3)
  
  se (n1 >=0 e n1< 10 e n2 >= 0 e n2 < 10 e n3 >=0 e n3 <10) {
  numComposicao = (n1* 100 + n2 *10 + n3)
   potencia = Matematica.potencia (numComposicao, 2.0)

   escreva ("O número resultado da composição é:", numComposicao )
   escreva ("\n")
   escreva ("O resultado do número elevado ao quadrado é:", potencia)}
      senao {escreva ("DIGITO INVALIDO")
      }
     
  }
}
