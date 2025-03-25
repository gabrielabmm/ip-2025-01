programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
inteiro numero, numeroQuad, i, ordemN, numeroB, soma=0

  escreva ("Qual é o número inteiro N?\n")
    leia (numero)

numeroB = numero/2

  para (i=1; i <= numeroB; i++)
  {soma = soma + 1
  ordemN = (soma * 2)
  numeroQuad = mat.potencia (ordemN, 2.0)
  escreva (ordemN, "^2 = ", numeroQuad, "\n")
  }
  }
}
