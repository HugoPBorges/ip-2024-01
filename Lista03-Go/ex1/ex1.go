package main

import "fmt"

func main() {
	var numero, teste int

	fmt.Scanln(&numero)
	numseq := make([]int, numero)
	for i := 0; i < numero; i++ {
		fmt.Scanln(&numseq[i])
	}

	fmt.Scanln(&teste)
	testes := make([]int, teste)
	for j := 0; j < teste; j++ {
		fmt.Scan(&testes[j])
	}

	for _, num := range testes {
		encontrado := false
		for _, comparacao := range numseq {
			if num == comparacao {
				encontrado = true
				break
			}
		}
		if encontrado == true {
			fmt.Println("ACHEI")
		} else {
			fmt.Println("NÃO ACHEI")
		}
	}
}
