package main

import "fmt"

func main() {
	// создание множеств
	var set1 = map[int]int{1: 1, 2: 1, 5: 1, 3: 1, 9: 1}
	var set2 = map[int]int{1: 1, 2: 1, 6: 1, 3: 1, 8: 1}

	for key, _ := range set1 {
		if _, ok := set2[key]; ok { // если оба словаря имеют текущий ключ печатаем его
			fmt.Print(key, " ")
		}
	}
}
