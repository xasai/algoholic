package main

import (
	"fmt"
	"sort"
)

func main() {

	var BoxCount int

	fmt.Scanf("%d", &BoxCount)

	Weights := make([]int, BoxCount)

	for i := 0; i < BoxCount; i++ {
		fmt.Scanf("%d", &Weights[i])
	}
	sort.Ints(Weights)
	TotalToAdd := 0
	for i := 1; i < BoxCount; i += 2 {
		TotalToAdd += Weights[i] - Weights[i-1]
	}
	fmt.Println(TotalToAdd)
}
