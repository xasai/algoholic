package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 35000), 800000000)

	var FirstDay, LastDay, Today int
	var RecordsNum int

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d %d %d", &FirstDay, &LastDay, &Today, &RecordsNum)

	graph := make(map[int]int, LastDay-FirstDay+1)

	for i := 0; i < RecordsNum; i++ {
		var day, value int

		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &day, &value)
		graph[day] = value
	}

	if _, ok := graph[FirstDay]; !ok {
		initFirst := -1
		graph[FirstDay] = initFirst
	}

	for i := 0; i <= LastDay-FirstDay; i++ {
		CurrentDay := FirstDay + i
		if CurrentDay > Today {
			graph[CurrentDay] = -1
		} else if value, ok := graph[CurrentDay]; (!ok || value == -1) && i != 0 {
			graph[CurrentDay] = graph[CurrentDay-1]
		}
		fmt.Println(CurrentDay, graph[CurrentDay])
	}
}
