package main

import (
	"fmt"
)

func main() {
	var n, i, K int
	var s float64
	if n >= 0 || n <= 9 {
		fmt.Scanln(&n)
	} else {
		fmt.Println("Número", n, "Inválido")
	}

	fmt.Scanln(&i)
	fmt.Scanln(&K)
	fmt.Scanln(&s)

	fmt.Println("Tabuada de Soma:")
	for g := 0; g < K; g++ {
		G := float64(i + (g)*int(s))
		Total := float64(n) + G
		fmt.Printf("%.2f + %.2f = %.2f\n", float64(n), G, Total)
	}

	fmt.Println("Tabuada de Subtração:")
	for g := 0; g < K; g++ {
		G := float64(i + (g)*int(s))
		Total := float64(n) - G
		fmt.Printf("%.2f - %.2f = %.2f\n", float64(n), G, Total)
	}

	fmt.Println("Tabuada de Multiplicação:")
	for g := 0; g < K; g++ {
		G := float64(i + (g)*int(s))
		Total := float64(n) * G
		fmt.Printf("%.2f x %.2f = %.2f\n", float64(n), G, Total)
	}

	fmt.Println("Tabuada de Divisão:")
	for g := 0; g < K; g++ {
		G := float64(i + (g)*int(s))
		Total := float64(n) / G
		fmt.Printf("%.2f / %.2f = %.2f\n", float64(n), G, Total)
	}
}
