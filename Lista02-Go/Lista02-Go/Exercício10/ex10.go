package main

import (
	"fmt"
)

func main() {
	var matricula, horas int
	var valor float64

	for {
		fmt.Scanln(&matricula, &horas, &valor)
		if matricula == 0 || horas == 0 || valor == 0.0 {
			break
		} else {
			Total := valor * float64(horas)
			fmt.Println(matricula, Total)
		}

	}

}
