package main

import "fmt" 

func main() {
	var numFatorial, i, num int64

  fmt.Println ("Digite um número inteiro")
  fmt.Scan (num)
  for i=0; i < num; i++ {
  numFatorial +=  num * (num - i)}

  fmt.Println(num,"! = ", numFatorial) 
}