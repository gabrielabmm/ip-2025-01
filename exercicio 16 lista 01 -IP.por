programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
  real salario, salarioReaj

    escreva ("Qual é o salário do funcionário? \n")
    leia (salario)

    se (salario < 300)
    salarioReaj = (salario * 1.5)

    senao salarioReaj = (salario * 1.3)

    escreva ("SALÁRIO COM REAJUSTE = ", mat.arredondar (salarioReaj, 2), "\n")
  }
}
