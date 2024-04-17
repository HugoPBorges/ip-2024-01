package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a, Ab, v float64
	fmt.Scanln(&h, &a)

	Ab = (3 * (a * a) * math.Sqrt(3)) / 2

	v = (Ab * h) / 3

	fmt.Println("O volume da piramide é :", math.Round(v*100)/(100), "metros cúbicos")

}
