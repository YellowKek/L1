package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Введите строку")
	var s string
	fmt.Scan(&s)
	fmt.Println(checkUnique(s))
}

func checkUnique(s string) bool {
	m := make(map[rune]int, len(s))
	strings.ToLower(s) // приведение к нижнему регистру
	for _, i := range s {
		_, ok := m[i]
		if ok { // если такой ключ уже встречался, значит такой символ уже был
			return false
		}
		m[i] = 1 // добавляем ключ
	}
	return true
}
