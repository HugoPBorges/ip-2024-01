package main

import (
	"fmt"
	"math"
)

func main() {
	var alunos [999]int
	var p1, p2, p3, p4, p5, p6, p7, p8 [999]float64
	var l1, l2, l3, l4, l5 [999]float64
	var trabalhofinal, presenca, NF, MP, ML [999]float64

	for i := 0; i < 999; i++ {

		fmt.Scanln(&alunos[i], &p1[i], &p2[i], &p3[i], &p4[i], &p5[i], &p6[i], &p7[i], &p8[i], &l1[i], &l2[i], &l3[i], &l4[i], &l5[i], &trabalhofinal[i], &presenca[i])
		if alunos[i] == 0 {
			break
		}

		MP[i] = (p1[i] + p2[i] + p3[i] + p4[i] + p5[i] + p6[i] + p7[i] + p8[i]) / 8

		ML[i] = (l1[i] + l2[i] + l3[i] + l4[i] + l5[i]) / 5

		NF[i] = 0.7*MP[i] + 0.15*ML[i] + 0.15*trabalhofinal[i]

		if presenca[i] > 96 && NF[i] >= 6 {
			s := "aprovado"
			fmt.Println("Matricula:", alunos[i], " Nota Final:", math.Round(NF[i]*100)/(100), "Situacao Final:", s)
		}
		if NF[i] < 6 {
			s := "Reprovado por nota"
			fmt.Println("Matricula:", alunos[i], " Nota Final:", math.Round(NF[i]*100)/(100), "Situacao Final:", s)
		}
		if presenca[i] < 96 {
			s := "Reprovado por frequência"
			fmt.Println("Matricula:", alunos[i], " Nota Final:", math.Round(NF[i]*100)/(100), "Situacao Final:", s)
		}
		if presenca[i] < 96 && NF[i] < 6 {
			s := "Reprovado por nota e frequência"
			fmt.Println("Matricula:", alunos[i], " Nota Final:", math.Round(NF[i]*100)/(100), "Situacao Final:", s)
		}

	}
}
