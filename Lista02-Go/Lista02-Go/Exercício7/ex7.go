package main

import (
	"fmt"
)

func main() {
	var (
		N, somaP, somaI, NP, NI int
	)
	NP = 0
	NI = 0

	for {
		fmt.Scanln(N)
		if N == 0 {
			break
		}
		if N%2 == 0 {
			somaP += N
			NP++

		} else {
			somaI += N
			NI++
		}

		mediaP := somaP / int(NP)
		mediaI := somaI / int(NI)

		fmt.Println("Média Par =", mediaP)
		fmt.Println("Média Impar =", mediaI)

	}

}
