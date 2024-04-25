package main

import (
	"fmt"
)

func main() {
	var ValorIngresso, ValorInicial, ValorFinal, V, N, Lucro float64

	fmt.Scanln(&ValorIngresso, &ValorInicial, &ValorFinal)

	if ValorFinal < ValorInicial {
		return
	}

	for x := 0; x < (ValorInicial-ValorFinal)+1; x++ {

		V[x] = ValorInicial + (float64(x))
		if ValorIngresso > V[x] {

			N[x] = 120 + (ValorIngresso-V[x])*50
		} else {
			N[x] = 120 - (V[x]-ValorIngresso)*60
		}
		Lucro[x] = N[x]*V[x] - (200 + 0.05*N[x])

		fmt.Println("V:", V[x], "N:", N[x], "L:", Lucro[x])
	}

	fmt.Println("Tabuada de Soma:")
	for g := 0; g < K; g++ {
		G := float64(i + (g)*int(s))
		Total := float64(n) + G
		fmt.Printf("%.2f + %.2f = %.2f\n", float64(n), G, Total)
	}
}
