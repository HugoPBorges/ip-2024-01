package main

import (
	"fmt"
)

func main() {

	var numero, total int

	fmt.Scan(&numero)
	fmt.Println(numero, " = 1")
	for i := 1; i < numero; i++ {
		if numero%i == 0 {
			total += i
			if i != 1 {

				fmt.Print("+", i)
			}
		}
	}
	if total == numero {

		fmt.Println(" =", total, "Numero perfeito")
	} else {
		fmt.Println(" =", total, "Numero não perfeito")
	}

}
