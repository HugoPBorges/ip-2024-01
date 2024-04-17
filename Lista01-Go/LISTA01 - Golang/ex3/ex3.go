package main

import "fmt"

func main() {
	var n1, n2, n3 int

	fmt.Scanln(&n1, &n2, &n3)

	var numero int = n1*100 + n2*10 + n3

	if n1 < 10 && n2 < 10 && n3 < 10 {
		fmt.Println(numero, numero*numero)

	} else {
		fmt.Println("Dígito Inválido")
	}

}
