package main

import "fmt"

func main() {
	for {
		var Num, maiornumero, indice, i int

		fmt.Scan(&Num)

		if Num == 0 {
			break
		}

		numseq := make([]int, Num)

		for i := 0; i < Num; i++ {
			fmt.Scan(&numseq[i])
			if numseq[i] > maiornumero {
				maiornumero = numseq[i]
				indice = i + 1
			}
		}
		i++
		for g := 0; g < i; g++ {
			fmt.Printf("índice:%d, maior numero:%d\n", indice, maiornumero)

		}
	}
}
