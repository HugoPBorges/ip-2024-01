package main

import "fmt"

func main() {
	var numero, ultimo, maiornum int
	posição := 0
	contagem := 0
	fmt.Scanln(&numero)
	numseq := make([]int, numero)

	for i := 0; i < numero; i++ {
		fmt.Scanln(&numseq[i])
		if i == numero-1 {
			ultimo = numseq[i]
		}
	}

	for i, teste := range numseq {
		if teste > maiornum {
			maiornum = teste
			posição = i + 1
		}
	}

	for _, num := range numseq {
		if num == ultimo {
			contagem++
		}
	}

	fmt.Printf("Nota %d, %d vezes\n", ultimo, contagem)
	fmt.Printf("Nota %d, indice %d", maiornum, posição)

}
