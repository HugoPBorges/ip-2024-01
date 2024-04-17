package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Scanln(&a)

	if a%3 == 0 && a%5 == 0 {
		fmt.Println("O número é divisível")
	} else {
		fmt.Println("O número não é divisível")
	}

}
