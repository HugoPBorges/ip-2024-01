package main

import "fmt"

func main() {
	var n int
	menor := 100000
	maior := 0

	fmt.Scanln(&n)
	seq := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&seq[i])
		fmt.Print(seq[i], " ")

		if seq[i] < menor {
			menor = seq[i]
		}
		if seq[i] > maior {
			maior = seq[i]
		}
	}

	fmt.Println()

	for j := n - 1; j >= 0; j-- {
		fmt.Print(seq[j], " ")

	}
	fmt.Println()
	fmt.Println(maior)
	fmt.Println(menor)
}
