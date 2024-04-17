package main

import (
	"fmt"
)

func main() {
	var nota float64
	fmt.Scanln(&nota)

	if 9.0 <= nota && nota <= 10.0 {
		fmt.Println("Nota:", nota, "Conceito = A")
	}

	if 7.5 <= nota && nota < 9.0 {
		fmt.Println("Nota:", nota, "Conceito = B")
	}

	if 6.0 <= nota && nota < 7.5 {
		fmt.Println("Nota:", nota, "Conceito = C")
	}

	if 0.0 <= nota && nota < 6.0 {
		fmt.Println("Nota:", nota, "Conceito = D")
	}

}
