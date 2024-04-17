package main

import (
	"fmt"
)

func main() {
	var numero, z int
	fmt.Scanln(&numero)

	z = numero%3*5 + numero/3*10

	fmt.Println("O valor a pagar é:", z)

}
