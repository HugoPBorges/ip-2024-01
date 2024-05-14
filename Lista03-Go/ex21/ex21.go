package main

import (
	"fmt"
	"sort"
)

func main() {
	var num int
	par := []int{}
	impar := []int{}

	fmt.Scanln(&num)
	seqnum := make([]int, num)

	for i := 0; i < num; i++ {
		fmt.Scanln(&seqnum[i])
	}
	for _, teste := range seqnum {

		if teste%2 == 0 {
			par = append(par, teste)

		} else {
			impar = append(impar, teste)
		}

	}
	sort.Ints(par)
	sort.Sort(sort.Reverse(sort.IntSlice(impar)))
	fmt.Println(par, "|", impar)

}
