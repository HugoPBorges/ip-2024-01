package main

import (
	"fmt"
)

func main() {

	var x, y int

	fmt.Scanln(&x, &y)

	if x%2 == 0 {
		fmt.Println(x)

		for i := 2; i <= y; i++ {
			x += 2

			fmt.Println(" ", x)

		}
	} else {
		fmt.Println("O primeiro número n é par")
	}

}
