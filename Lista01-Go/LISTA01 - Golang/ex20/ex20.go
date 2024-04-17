package main

import (
	"fmt"
)

func main() {
	var h, m, s int
	fmt.Scanln(&h)
	fmt.Scanln(&m)
	fmt.Scanln(&s)

	T := (h * 3600) + (m * 60) + s

	fmt.Println("O tempo em segundos é:", T)

}
