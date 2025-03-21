programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
   real raio, altura, area, areaCirculo, areaLateral, areaTotal, custo

  const real pi = 3.14159

    escreva ("Qual o raio da lata?\n")
    leia (raio)

    escreva ("Qual é a altura da lata?\n")
    leia (altura)

    escreva ("\n")
   
    areaCirculo = (pi * raio * raio)
    areaLateral = (2 * pi * raio * altura)
    areaTotal = (2 * areaCirculo + areaLateral)
    custo = (100 * areaTotal)
    
    escreva ("O VALOR DO CUSTO É= ", mat.arredondar (custo, 2 ))

    escreva ("\n")

  }
}
