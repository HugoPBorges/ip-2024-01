package main

import (
	"fmt"
)

func main() {
	var n int
	_, x := fmt.Scanln(&n)
	if x != nil || n <= 1 {

		fmt.Println("Numero invalido!")
	}

	soma := calcularSoma(n)
	fmt.Printf("%.6f\n", soma)
}

func calcularSoma(n int) float64 {
	var soma float64
	for z := 1; z <= n; z++ {
		soma += 1.0 / float64(z)
	}
	return soma
}
