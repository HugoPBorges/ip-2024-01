package main

import (
	"fmt"
	"math"
)

func main() {

	var x float64
	var X, num, den int
	den = 1
	numero := 2

	fmt.Scanln(&x)

	X = int(math.Round(x * 100))
	for {
		if X == 0 {
			if X%numero != 0 {
				numero++
			} else {
				den *= numero
			}

			num = den * int(x)

			fmt.Println(num, "/", den)
		}
	}
}
