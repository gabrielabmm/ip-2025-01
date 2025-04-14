package main

import "fmt"

func main() {
	var x, i int
	fmt.Scan(&x)
	soma := 0
	for i = 1; i <= x; i++ {
		fmt.Println(i)
		soma += i
	}
	fmt.Println(soma)

}
