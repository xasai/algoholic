package main

import "fmt"

func main() {
	var FirstDay, LastDay, Today int
	var RecordsNum int

	fmt.Scanf("%d %d %d %d", &FirstDay, &LastDay, &Today, &RecordsNum)

	graph := make(map[int]int, LastDay-FirstDay+1)
	beforeFirst := make([]int, 0, 10)

	for i := 0; i < RecordsNum; i++ {
		var day, value int

		fmt.Scanf("%d %d", &day, &value)
		graph[day] = value
		if day < FirstDay {
			beforeFirst = append(beforeFirst, day)
		}
	}

	initFirst := -1

	for i := len(beforeFirst) - 1; i >= 0; i-- {
		CurrentDay := graph[beforeFirst[i]]
		if CurrentDay > Today {
			initFirst = -1
			break
		} else if value, ok := graph[CurrentDay]; (!ok || value == -1) && i != len(beforeFirst)-1 {
			graph[CurrentDay] = graph[beforeFirst[i-1]]
		}
	}

	graph[Firstday-1] = initFirst
	//fmt.Println("Firstday ", graph[FirstDay])

	for i := 0; i <= LastDay-FirstDay; i++ {
		CurrentDay := FirstDay + i
		if CurrentDay > Today {
			graph[CurrentDay] = -1
		} else if value, ok := graph[CurrentDay]; !ok || value == -1 {
			graph[CurrentDay] = graph[CurrentDay-1]
		}
		fmt.Println(CurrentDay, graph[CurrentDay])
	}
}
