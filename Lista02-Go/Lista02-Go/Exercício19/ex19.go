package main

import (
	"fmt"
)

func main() {

	var numero, result int

	fmt.Scan(&numero)

	start := 1

	for x := 1; x <= numero; x++ {

		fmt.Println(x, "*", x, "*", x, "=")

		result = start + (2 * (x - 1))

		for c := start; c <= result; c += 2 {

			fmt.Println(c)

			if c != result {

				fmt.Print("+")

			} else {

			}
		}
		start = result + 2

	}
}
