package main

import "fmt"

func main() {
	var numero, soma int

	fmt.Scanln(&numero)

	seqnum := make([]int, numero)

	for i := 0; i < numero; i++ {
		fmt.Scan(&seqnum[i])
		soma += seqnum[i]

	}
	fmt.Printf("Soma; %d", soma)
}
