programa {
  inclua biblioteca Matematica --> mat 
  funcao inicio() {
    real altura, aresta, volume, areaHex 

    escreva ("Qual é a altura da pirâmide? \n")
    leia (altura)

    escreva ("\n")

    escreva ("Qual número corresponde a aresta do hexágono que forma a base da pirâmide? \n")
    leia (aresta)

    areaHex = ((3 * mat.potencia (aresta, 2) * mat.raiz (3, 2)) /2)
    volume = ((areaHex * altura) /3)
    
   escreva ("O VOLUME DA PIRÂMIDE = ", mat.arredondar (volume, 2), " METROS CUBICOS")
  }
}
