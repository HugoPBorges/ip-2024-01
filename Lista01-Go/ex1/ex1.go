package main

import "fmt"

func main() {
	var nota1, nota2, nota3 int

	fmt.Scanln(&nota1, &nota2, &nota3)

	x := (nota1 + nota2 + nota3) / 3

	if x < 6 {
		fmt.Println("Média: ", x, "\n", "Reprovado")
	} else {
		fmt.Println("Média: ", x, "\n", "Aprovado")

	}

}
