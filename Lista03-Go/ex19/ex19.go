package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scanln(&num)

	elementosunicos := make(map[int]bool)

	for i := 0; i < num; i++ {
		var elementos string

		fmt.Scanln(&elementos)
		num, err := strconv.Atoi(elementos)
		if err != nil {
			continue
		}
		elementosunicos[num] = true
	}

	for sequencia := range elementosunicos {
		fmt.Println(sequencia)
	}
}
