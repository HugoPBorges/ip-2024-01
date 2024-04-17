package main

import (
	"fmt"
	"math"
)

func main() {
	var S, NS float64

	fmt.Scanln(&S)

	if S <= 300 {
		NS = S * .5
		fmt.Println("Salário com reajuste:", math.Round(NS*100)/(100))
	}
	if S > 300 {
		NS = S * 1.3
		fmt.Println("Salário com reajuste:", math.Round(NS*100)/(100))
	}
}
