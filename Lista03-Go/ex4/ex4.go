package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)
	seqnum := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&seqnum[i])
	}
	fmt.Println(seqnum)

}
