package main

import (
	"fmt"
)

func main() {

	var a, b, c int

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	Delta := (b * b) - 4*(a)*(c)

	fmt.Println("O valor de delta é : ", Delta)
}
