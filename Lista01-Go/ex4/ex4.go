package main

import "fmt"

func main() {
	var Sm float64
	var K float64

	fmt.Scanln(&Sm, &K)

	var Y float64 = 0.7 * Sm

	var uK float64 = Y / 100
	var C float64 = K * uK
	var CD float64 = C * 0.9

	fmt.Println("Custo por Kw:", uK)
	fmt.Println("Custo do consumo:", C)
	fmt.Println("Custo com desconto:", CD)

}
