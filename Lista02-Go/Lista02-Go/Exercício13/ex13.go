package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	Somafinal := 0 - n
	graos := n

	for i := 2; i < 66; i++ {
		Somafinal += graos
		if i%2 == 0 {
			graos *= 2
		} else {
			graos = n
		}
	}
	fmt.Println(Somafinal)
}
