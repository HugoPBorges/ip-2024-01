package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var mediana float64
	fmt.Scanln(&n)
	numseq := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&numseq[i])

	}
	sort.Ints(numseq)

	if n%2 == 0 {
		mediana = float64(numseq[(n/2)-1]+numseq[n/2]) / 2
	} else {
		mediana = float64(numseq[(n / 2)])
	}
	fmt.Printf("%.2f\n", mediana)
}
