package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type good struct {
	warehouse     int
	deliveryPrice int
}

var stdin *bufio.Scanner

func scanInt() int {
	stdin.Scan()
	ret, _ := strconv.ParseInt(stdin.Text(), 10, 32)
	return int(ret)
}

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	stdin = bufio.NewScanner(bytes.NewReader(b))
	stdin.Split(bufio.ScanWords)

	warehouseCount := scanInt()

	goods := make(map[int]good) //key = good's articul

	for i := 0; i < warehouseCount; i++ {

		warehouseNum := scanInt()
		goodsCount := scanInt()

		for k := 0; k < goodsCount; k++ {
			articul := scanInt()
			price := scanInt()
			newGood := good{
				warehouse:     warehouseNum,
				deliveryPrice: price,
			}

			g, ok := goods[articul]
			if !ok || g.deliveryPrice > newGood.deliveryPrice {
				goods[articul] = newGood
			} else if g.deliveryPrice == newGood.deliveryPrice && g.warehouse > newGood.warehouse {
				goods[articul] = newGood
			}
		}
	}

	ordersCount := scanInt()

	for i := 0; i < ordersCount; i++ {
		articul := scanInt()
		price := scanInt()
		if good, ok := goods[articul]; ok {
			_ = good
			fmt.Println(good.warehouse, good.deliveryPrice+price)
		} else {
			fmt.Println("-1 -1")
		}
	}
}
