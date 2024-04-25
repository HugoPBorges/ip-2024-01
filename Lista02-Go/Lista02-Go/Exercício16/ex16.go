package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	for h := 1; h <= n; h++ {
		for c1 := 1; c1 < h; c1++ {
			for c2 := c1 + 1; c2 < h; c2++ {

				if c1*c1+c2*c2 == h*h {
					fmt.Println("hipotenusa:", h, "Cateto1:", c1, "Cateto2:", c2)
				}
			}

		}

	}

}
