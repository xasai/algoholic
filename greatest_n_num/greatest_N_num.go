package main

import (
	"fmt"
	"sort"
)

func main() {
	var M, N int

	fmt.Scanf("%d %d", &M, &N)

	slice := make([]int, 0, M)
	occured := make(map[int]struct{})

	for i := 0; i < M; i++ {
		n := 0
		fmt.Scanf("%d", &n)
		if _, ok := occured[n]; !ok {
			occured[n] = struct{}{}
			slice = append(slice, n)
		}
	}

	sort.Ints(slice)

	if len(slice) >= N {
		fmt.Println(slice[len(slice)-N])
	} else {
		fmt.Println(-1)
	}
}
