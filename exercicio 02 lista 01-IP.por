programa {
  inclua biblioteca Matematica --> mat

  funcao inicio() { 
    inteiro numJogos, i
    real numPessoas, percPo, percGe, percAr, percCa, renda1, renda2, renda3, renda4, rendaTotal

    escreva ("Qual o número de casos de teste?\n")
    leia (numJogos)  
    escreva ("\n")

    para (i = 1; i <= numJogos; i++) {

      escreva ("Quantas pessoas compraram ingresso para o jogo ", i, "?\n")
      leia (numPessoas)
      
      escreva ("Qual a percentagem de pessoas que compraram na categoria Popular?\n")
      leia (percPo)

      escreva ("Qual a percentagem de pessoas que compraram na categoria Geral?\n")
      leia (percGe) 
      
      escreva ("Qual a percentagem de pessoas que compraram na categoria Arquibancada?\n")
      leia (percAr)

      escreva ("Qual a percentagem de pessoas que compraram na categoria Cadeiras?\n")
      leia (percCa)

      // Ajuste no cálculo da renda total
      renda1 = (numPessoas * percPo) / 100
      renda2 = (numPessoas * percGe * 5) / 100
      renda3 = (numPessoas * percAr * 10) / 100
      renda4 = (numPessoas * percCa * 20) / 100

      rendaTotal = renda1 + renda2 + renda3 + renda4

      
    }

    para (i = 1; i <= numJogos; i++) {
      escreva ("A renda do jogo ", i, " é ", rendaTotal, "\n\n")
    }
  }
}
