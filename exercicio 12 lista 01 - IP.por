programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    inteiro horas, valor, preco
   
    

  escreva ("Por quantas horas a charrete foi utilizada? \n")
  leia (horas)

preco = 5 * horas

 se (horas % 3 == 0) {
  valor = preco - 5 }

 senao {valor = preco}

 escreva ("O VALOR A PAGAR É = ", valor, "\n")
  
    
  }
}
