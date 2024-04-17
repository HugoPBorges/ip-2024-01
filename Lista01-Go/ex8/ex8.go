package main

import (
	"fmt"
	"math"
)

func main() {
	var R, h, Ac, Al, At, x float64
	fmt.Scanln(&R)
	fmt.Scanln(&h)
	var π float64 = 3.14

	Ac = π * (R * R)

	Al = 2 * π * (R) * (h)

	At = 2*Ac + Al

	x = 100 * At

	fmt.Println("O valor de custo é:", math.Round(x*100)/(100))

}
