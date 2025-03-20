programa 
{ inclua biblioteca Matematica --> Mat

  funcao inicio() {
    real qntdTemp, i, Celsius,qntdTemp, t, conversao, conversao2

    escreva ("Quantas temperaturas irão ser convertidas de Fahrenheit para Celsius? \n")
    leia (qntdTemp)

    para (i=1; i<= qntdTemp; i ++) {

      escreva ("Qual é a temperatura ", i, "?\n")
      leia (t)

      conversao = ((5 * (t - 32)) /9)

      escreva (t ," FAHRENHEIT EQUIVALE A ", Mat.arredondar (conversao, 2), " CELSIUS \n")
     
    }
  }
}
