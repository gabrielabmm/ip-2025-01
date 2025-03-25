programa {
  funcao inicio() {
    real horas, min, seg, tempo
   
    escreva ("Qual o valor do tempo expresso em horas?\n")
    leia (horas)

     escreva ("Qual o valor do tempo expresso em minutos?\n")
    leia (min)

     escreva ("Qual o valor do tempo expresso em segundos?\n")
    leia (seg)

    tempo = (horas * 3600 + min* 60 + seg)
    
    escreva ("O TEMPO EM SEGUNDOS É = ", tempo, "\n")
  }
}
