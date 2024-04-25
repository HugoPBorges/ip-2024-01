package main

import (
	"fmt"
)

func main() {
	var numero int

	fmt.Scanln(&numero)

	if numero < 1 {
		fmt.Println("Número não Primo")
	} else if numero%2 == 0 || numero%3 == 0 {
		fmt.Println("Número não Primo")
	} else {
		fmt.Println("Número Primo")
	}

}
