programa {
  funcao inicio() {
  real valorIn, razao, numElem, i, soma=0

    escreva ("Qual o valor inicial da progressão aritmética? \n")
    leia (valorIn)

    escreva ("Qual a razão da progressão? \n")
    leia (razao)

    escreva ("Quantos elementos formam a progressão? \n")
    leia (numElem)

    para (i=0; i < numElem; i++)
    soma = valorIn + razao * i + soma

  escreva ("O valor da soma dos ", numElem, " elementos é = ", soma)

    
  }
}
