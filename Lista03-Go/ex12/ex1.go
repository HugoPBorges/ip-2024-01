package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	seq := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&seq[i])

	}
	sort.Ints(seq)
	fmt.Println(seq)
}
