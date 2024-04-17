package main

import (
	"fmt"
)

func main() {

	var Vi, R, n int

	Vf := (2*n*Vi + n*n*R - 1*n*R) / 2

	fmt.Scanln(&Vi, &R, &n)

	fmt.Sprintln("Soma:", Vf)

}
