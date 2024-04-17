package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Scanln(&n)

	y := make([]float64, n, n)

	for i := 0; i < n; i++ {

		fmt.Scanln(&y[i])
	}

	for i := 0; i < n; i++ {

		C := (5*y[i] - 160) / 9
		fmt.Println(y[i], "Fahrenheit equivale a ", math.Round(C*100)/(100), "Celsius")

	}

}
