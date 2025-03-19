programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() { 
    real numJogos, numPessoas, percPo, percGe, percAr, percCa, renda1, renda2, renda3, renda4, rendaTotal
    escreva ("Qual o número de casos de teste?\n")
    leia (numJogos)
    escreva ("\n")

    escreva ("Quantas pessoas compraram ingresso pro jogo x?\n")
    leia (numPessoas)
    
    escreva ("Qual a percentagem de pessoas que compraram na categoria Popular?\n")
    leia (percPo)

     escreva ("Qual a percentagem de pessoas que compraram na categoria Geral?\n")
    leia (percGe) 
    
    escreva ("Qual a percentagem de pessoas que compraram na categoria Arquibancada? \n")
    leia (percAr)

     escreva ("Qual a percentagem de pessoas que compraram na categoria Cadeiras? \n")
    leia (percCa)

    renda1 = (numPessoas * percPo)
    renda2 = (numPessoas * percGe * 5)
    renda3 = (numPessoas * percAr * 10)
    renda4 = (numPessoas * percCa * 20) 

    rendaTotal = ((renda1 + renda2 + renda3 + renda4) /100)

    escreva ("A renda do jogo x é \n", rendaTotal)
  }
} 
