package main

import (
	"bufio"
	"fmt"
	"os"
)

type heap struct {
	max int
	min int
}

func search(n int, heaps []heap) int {
	low := 0
	high := len(heaps) - 1

	for low <= high {
		mid := (low + high) / 2
		heap := heaps[mid]
		if heap.max >= n && heap.min <= n {
			return mid + 1
		}
		if heap.min > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func intScanln(n int, s string) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Sscanln(s, y...)
	x = x[:n]
	return x, err
}

func main() {

	buf := make([]byte, 0, 32768)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(buf, 4294967296)

	heapCount := 0
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &heapCount)

	heaps := make([]heap, 0, heapCount)
	scanner.Scan()
	heapSizes, _ := intScanln(heapCount, scanner.Text())

	min := 0
	max := 0

	for i := 0; i < heapCount; i++ {
		heapSize := heapSizes[i]
		min = max + 1
		max += heapSize
		//fmt.Printf("Heap [%d] max: %d min: %d\n", i+1, min, max)
		heaps = append(heaps, heap{max: max, min: min})
	}

	productCount := 0
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &productCount)
	//fmt.Printf("product count: %d\n", productCount)

	answers := make([]int, 0, productCount)

	scanner.Scan()
	products, _ := intScanln(productCount, scanner.Text())

	for i := 0; i < productCount; i++ {
		//fmt.Printf("Searching %d : %d\n", productNum, search(productNum, heaps))
		answers = append(answers, search(products[i], heaps))
	}
	for _, answ := range answers {
		fmt.Println(answ)
	}
}
