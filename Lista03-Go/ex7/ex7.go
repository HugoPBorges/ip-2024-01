package main

import (
	"fmt"
)

func main() {
	var N int
	maiornum := 0
	contagem := make([]int, maiornum+1)
	contador := 0

	for {
		fmt.Scanln(&N)
		if N == 0 {
			break
		}

		seq := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&seq[i])
		}

		for _, num := range seq {
			if num > maiornum {
				maiornum = num
			}
		}
		for j := 0; j < maiornum; j++ {
			for _, num := range seq {
				if num <= j {
					contador++
				}
			}

			for j, contador = range contagem {
				fmt.Printf("(%d) %d\n", j, contador)
			}
		}

	}
}
