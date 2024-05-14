package main

import "fmt"

func main() {
	var ntestes, teste int

	fmt.Scanln(&ntestes)
	for i := 0; i < ntestes; i++ {
		fmt.Scanln(&teste)
		seqnum := make([]int, teste)
		for i := 0; i < teste; i++ {
			fmt.Scan(&seqnum[i])
			fmt.Print(seqnum[i], " ")
		}

	}

}
