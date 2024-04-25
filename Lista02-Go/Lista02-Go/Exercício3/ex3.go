package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scanln(&n)

	fatorial := 1
	for i := 0; i < n; i++ {
		fatorial *= n - i
	}
	fmt.Println(n, "! = ", fatorial)
}
