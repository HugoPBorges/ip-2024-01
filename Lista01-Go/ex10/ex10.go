package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, Det int

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)

	Det = (a)*(d) - (b)*(c)

	fmt.Println("O valor do determinante é :", Det)
}
