package main

import (
	"fmt"
)

func main() {
	var n, contagem, maior int
	fmt.Scanln(&n)

	numseq := make([]int, n)
	freq := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&numseq[i])
		if numseq[i] > 0 && numseq[i] < 100 {
			for j := 0; j < n; j++ {

				if numseq[i] == numseq[j] {

					freq[i]++
					for _, contagem := range freq {
						if contagem > maior {
							maior = contagem

						}
					}
				}
			}
		}
	}
	fmt.Println(contagem)
}
