package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numero int
	for {

		fmt.Scanln(&numero)

		numbin := strconv.FormatInt(int64(numero), 2)
		fmt.Println(numbin)

	}
}
