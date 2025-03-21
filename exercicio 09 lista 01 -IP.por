programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real coefA, coefB, coefC, delta

    escreva ("Quais os coeficientes A, B e C?\n")
  leia (coefA)
  leia (coefB)
  leia (coefC)

  delta = (mat.potencia (coefB, 2) - 4 * coefA * coefC)

  escreva ("O VALOR DE DELTA É = ", mat.arredondar (delta, 2))


  }
}
