package main

import "fmt"

func somaD(n int) int {
	amigos := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			amigos += i
			if n/i != i {
				amigos += n / i
			}
		}
	}
	return amigos
}

func main() {
	var num int
	fmt.Scan(&num)

	contagem := 0
	for i := 2; contagem < num; i++ {

		amigos := somaD(i)
		if amigos > i && somaD(amigos) == i {
			fmt.Printf("(%d,%d)\n", i, amigos)
			contagem++
		}
	}
}
