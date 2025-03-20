programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
real contaAgua, consumo, residencial, comercial, industrial
caracter c, C, R, r, I, i, tipo
    escreva ("Qual o valor da conta de água? \n ")
    leia (contaAgua)

    escreva ("Qual é o consumo de água por metros cúbicos? \n")
    leia (consumo)

    escreva ("Qual o tipo de consumidor? (C - Comercial, R - Residencial, I - Industrial) \n)")
    leia (tipo)

    escreva ("CONTA=", contaAgua)
    escreva ("\n")

  residencial = (5 + 0.05 * contaAgua)
  comercial = (500 + 0.25 * (contaAgua - 80))
  industrial = (800 + 0.04 * (contaAgua - 100))

  escolha (tipo)
   {   caso "C":
    escreva ("VALOR DA CONTA=", comercial)
pare
   
   caso "R":
    escreva ("VALOR DA CONTA =", residencial )
   pare

   caso "I":
    escreva ("VALOR DA CONTA =", industrial)
pare
   }
  }
}
