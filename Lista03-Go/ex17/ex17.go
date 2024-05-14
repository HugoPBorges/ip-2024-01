package main

import "fmt"

func main() {
	var n, contagem int
	fmt.Scanln(&n)
	numseq := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numseq[i])
	}

	for i := 0; i < n; i++ {
		valoratual := numseq[i]
		unico := false

		for j := 0; j < n; j++ {
			if i != j && valoratual == numseq[j] {
				unico = true
				break
			}
		}

		if !unico {
			contagem++
		}
	}

	fmt.Println(contagem)
}
