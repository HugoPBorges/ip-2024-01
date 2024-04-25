package main

import (
	"fmt"
)

func main() {
	var n, seq, max int
	var y []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&seq)
		y = append(y, seq)

	}
	for i := 0; i < n-1; i++ {
		cont := 1
		for k := i + 1; k < n; k++ {
			if y[i] == y[k] {

				cont++

			} else {

				break
			}
		}
		if cont >= max {

			max = cont
		}
	}
	fmt.Println("A maior subsequencia de elementos iguais possui", max, "elementos")
}
