package main

import (
	"fmt"
)

var bool2int = map[bool]int{false: 0, true: 1}

func main() {

	var LenAvailable int
	var NeedLen int

	fmt.Scanf("%d %d", &LenAvailable, &NeedLen)

	if LenAvailable == 0 || NeedLen == 0 {
		fmt.Println(0)
		return
	}
	res := NeedLen/LenAvailable + bool2int[NeedLen%LenAvailable != 0]
	fmt.Println(res)
}
