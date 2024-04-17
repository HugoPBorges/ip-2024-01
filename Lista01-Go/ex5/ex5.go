package main

import "fmt"

func main() {
	var Cc int
	var CA, Custo float64
	var T string

	fmt.Scanln(&Cc, &CA, &T)
	if T == "R" {
		Custo = 5 + 0.05*CA
	}

	if T == "C" {
		if CA < 80 {
			Custo = 500
		} else {
			Custo = 500 + 0.25*CA
		}
	}

	if T == "I" {
		if CA < 100 {
			Custo = 800
		} else {
			Custo = 800 + 0.04*CA
		}
	}

	fmt.Println("Conta:", Cc)
	fmt.Println("Valor da conta:", Custo)

}
