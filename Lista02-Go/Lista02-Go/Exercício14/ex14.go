package main

import "fmt"

func main() {
	var linha, coluna int
	fmt.Scan(&linha, &coluna)

	for i := 0; i < linha; i++ {
		for j := 0; j < coluna; j++ {
			if j < i {
				fmt.Printf("(%d-%d) ", i, j)
			}
		}
		fmt.Println()
	}
}
