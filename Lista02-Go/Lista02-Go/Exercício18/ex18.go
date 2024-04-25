package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	número := 2
	mmc := 1
	for {
		if a == 1 && b == 1 && c == 1 {
			break
		}
		if a%número != 0 && b%número != 0 && c%número != 0 {
			número++
		} else {
			mmc *= número
			fmt.Println(a, b, c, "|", número)
			if a%número == 0 {
				a = a / número
			}
			if b%número == 0 {
				b = b / número
			}
			if c%número == 0 {
				c = c / número
			}
		}
	}
	fmt.Println("MMC :", mmc)
}
