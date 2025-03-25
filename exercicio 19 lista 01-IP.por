programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
inteiro numero, i
real soma = 0


    escreva ("Escreva um número inteiro positivo e maior que 1. \n")
    leia (numero)
    
    se (numero > 1) {
    para (i=0; i <= numero; i++) 
    soma = (1.0 / (i+1)) + soma 
    
   escreva (mat.arredondar (soma, 6)) }
    
    senao {
      escreva ("Número inválido!")}
    
    
  }
}
