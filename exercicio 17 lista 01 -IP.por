programa {
  funcao inicio() {
  inteiro numeroX, numeroY, i, soma

    escreva ("Quais os números x e y?\n")
    leia(numeroX)
    escreva ("\n")
    leia (numeroY)
    escreva ("\n")

    soma = numeroX - 2

    para (i= 0; i <= i+1; i++ ) {
  se (numeroX % 2 == 0) {
  soma = soma + 2
  escreva (soma, "\t")}

    senao {
    escreva ("O PRIMEIRO NÚMERO NÃO É PAR")
    pare 
    }


    }
    
  }
}
