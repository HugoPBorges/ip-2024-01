package main

import "fmt"

func main() {

	var N int
	fmt.Scanln(&N)

	for x := 2; x <= N; x += 2 {

		var z int

		z = x * x

		fmt.Println(x, "^2 =", z)

	}

}
