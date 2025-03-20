programa {
  inclua biblioteca Matematica --> Mat
  funcao inicio() {
    real valorF, valorC, conversaoFC, conversaoPM

    escreva ("Qual o valor, em Fahrenheit, que será convertido para Celsius? \n")
    leia (valorF)

    escreva ("Qual é a quantidade de chuva, em polegadas, que será convertida para milímetros? \n")
    leia (valorC)

    conversaoFC = ((5 * valorF - 160) /9)
    conversaoPM = (valorC * 25.4)

    escreva ("O VALOR EM CELSIUS = ", Mat.arredondar (conversaoFC, 2) , "\n")
    escreva ("A QUANTIDADE DE CHUVA É = ", Mat.arredondar (conversaoPM, 2))
  }
}
