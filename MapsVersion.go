package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(SortMyString("asdfJKGGSfbbbjjsiSDodwkloчтооооАААаааа"))
}

func SortMyString(str string) string {

arr := []rune(str)
var s string
result := make(map[rune]int)
	for _, v := range arr {
		result[v]++
	}

// сортировка ключей в порядке возрастания
keys := make([]rune, 0, len(result))
	for k := range result {
		keys = append(keys, k)
	}
sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	// форматированный вывод данных
	for _, k := range keys {
		s += fmt.Sprintf("%c%d", k, result[k])
	}
	return s
}
