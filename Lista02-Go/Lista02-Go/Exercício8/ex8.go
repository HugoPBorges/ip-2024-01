package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)

	if N <= 2 {
		fmt.Println("Número Inválido")
		return
	}

	finais := 1
	for i := 1; i <= N; i++ {
		for k := i + 1; k <= N; k++ {
			fmt.Println("Final", finais, ": Time", i, "X", "Time", k)
			finais++
		}
	}

}
