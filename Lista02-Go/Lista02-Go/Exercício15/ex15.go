package main

import (
	"fmt"
)

func main() {
	var T, A, B, Final1, Final2 int

	fmt.Scanln(&T)
	for i := 0; i < T; i++ {
		fmt.Scanln(&A, &B)

		XA := A % 10
		YA := (A%100 - A%10) / 10
		ZA := (A - A%100 - A%10) / 100
		Final1 = XA*100 + YA*10 + ZA

		XB := B % 10
		YB := (B%100 - B%10) / 10
		ZB := (B - B%100 - B%10) / 100
		Final2 = XB*100 + YB*10 + ZB

		if Final1 > Final2 {
			fmt.Println(Final1)
		} else {
			fmt.Println(Final2)
		}
	}

}
