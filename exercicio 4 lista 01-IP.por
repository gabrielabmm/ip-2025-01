programa {
  funcao inicio() {
    
  real salarioMin, kwGastos, kWunico, valorkW, valorRes, novoValor

    escreva ("Qual é o valor do salário mínimo? \n")
    leia (salarioMin)

    escreva ("Quantos kW foram gastos pela residência? \n")
    leia (kwGastos)

  kWunico = (salarioMin * 7/10)
  valorkW = (kWunico /100)

   escreva ("O valor, em real, de cada kW é ", valorkW)
   escreva ("\n")

   valorRes = (valorkW * kwGastos)

   escreva ("O valor, em reais, a ser pago pela residência é:", valorRes )
   escreva ("\n")
  novoValor = (valorRes * 9/10)

   escreva ("O novo valor a ser pago, com desconto de 10%, é:", novoValor)
  }

}

