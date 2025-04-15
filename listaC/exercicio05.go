package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	if Triangular(num) {
		fmt.Println(num, "eh triangular")
	} else {
		fmt.Println(num, "nao eh triangular")
	}
}

func Triangular(num int) bool {
	for i := 1; i*(i+1)*(i+2) <= num; i++ {
		if i*(i+1)*(i+2) == num {
			return true
		}
	}
	return false
}
