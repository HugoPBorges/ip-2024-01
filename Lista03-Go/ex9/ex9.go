package main

import (
	"fmt"
	"math"
)

func distancia(xA, yA, zA, xB, yB, zB float64) float64 {

	return math.Sqrt(math.Pow(xA-xB, 2) + math.Pow(yA-yB, 2) + math.Pow(zA-zB, 2))
}

func main() {
	var vezes int
	fmt.Scanln(&vezes)

	cordenadas := make([]struct{ x, y, z float64 }, vezes)
	for i := 0; i < vezes; i++ {
		fmt.Scanln(&cordenadas[i].x, &cordenadas[i].y, &cordenadas[i].z)
	}

	for i := 1; i < vezes; i++ {
		dist := distancia(cordenadas[i-1].x, cordenadas[i-1].y, cordenadas[i-1].z, cordenadas[i].x, cordenadas[i].y, cordenadas[i].z)
		fmt.Printf("%.2f\n", dist)
	}
}
