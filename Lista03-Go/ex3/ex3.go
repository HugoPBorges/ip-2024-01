package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)
	seq := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&seq[i])

	}

	for j := n - 1; j >= 0; j-- {
		fmt.Print(seq[j], " ")
	}
}
