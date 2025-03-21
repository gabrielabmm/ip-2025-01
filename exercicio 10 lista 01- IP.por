programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
   real elementoA, elementoB, elementoC, elementoD, det

escreva ("Quais os quatro elementos (a,b,c e d) da matriz? \n")
leia (elementoA)
leia (elementoB)
leia (elementoC)
leia (elementoD)

det = ((elementoA * elementoD) - (elementoB * elementoC))

escreva ("O VALOR DO DETERMINANTE É = ", mat.arredondar (det, 2))
    
  }
}
