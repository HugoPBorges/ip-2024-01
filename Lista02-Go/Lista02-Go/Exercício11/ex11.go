package main

import (
	"fmt"
	"math"
)

func main() {

	var num, erro float64

	fmt.Scan(&num, &erro)

	r := 1.0

	for math.Abs(num-r*r) > erro {

		r = (r + num/r) / 2

		fmt.Println("r:", r, "erro:", math.Abs(num-r*r))
	}
}
