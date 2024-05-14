package main

import "fmt"

func main() {
	var testes int
	fmt.Scanln(&testes)

	soma := make([]int, testes)

	for i := 0; i < testes; i++ {
		soma[i] = 0
		digitos := make([]int, 11)

		for j := 0; j < 11; j++ {
			fmt.Scan(&digitos[j])
			soma[i] += digitos[j]

		}
	}
	for i := 0; i < testes; i++ {
		if soma[i]%11 == 0 {
			fmt.Println("CPF válido")
		} else {
			fmt.Println("CPF inválido")
		}

	}

}
