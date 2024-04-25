package main

import (
	"fmt"
)

func main() {
	var A, B float64
	fmt.Scanln(&A)
	fmt.Scanln(&B)

	anos := 0
	for A < B {
		A = A * 1.03
		B = B * 1.015
		anos++
	}
	fmt.Println("Anos =", anos)
}
