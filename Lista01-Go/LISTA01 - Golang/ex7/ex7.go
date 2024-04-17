package main

import (
	"fmt"
	"math"
)

func main() {
	var Fa, Po, x, y float64
	fmt.Scanln(&Fa)
	fmt.Scanln(&Po)
	x = (5*Fa - 160) / 9

	y = Po * 25.4
	fmt.Println("O valor em celsius=", math.Round(x*100)/(100), "\n", "A quantidade de chuva é =", math.Round(y*100)/(100))

}
