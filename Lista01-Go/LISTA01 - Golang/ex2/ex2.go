package main

import "fmt"

func main() {
	var NT int // número de testes
	var NP float64
	var PP, PG, PA, PC, x float64 // porcentagem popular, geral, arquibancada e cadeiras
	//  PP = p / PG = g / PA = a / PC = c

	fmt.Scanln(&NT)
	n := make([]float64, NT, NT)

	for i := 0; i <= NT; i++ {

		fmt.Scanln(&NP[i], &PP[i], &PG[i], &PA[i], &PC[i])

	}

	for i := 0; i <= NT; i++ {

		x = PP[i]*NP[i]*1.0 + PG[i]*NP[i]*5.0 + PA[i]*NP[i]*10.0 + PC[i]*NP[i]*20.0

		fmt.Println("A renda do jogo", n[i], "é:", (x*100)/(100))
	}
}
