package main

import "fmt"

func main() {
	var N, num, contagem int

	fmt.Scanln(&N)

	if N >= 1 && N <= 1000 {
		V := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&V[i])
		}

		fmt.Scanln(&num)

		for _, valor := range V {
			if valor >= num {
				contagem++
			}
		}

		fmt.Println(contagem)
	} else {
		fmt.Println("digite um número valido")
	}
}
