package main

import "fmt"

func main() {
	var input = []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]int) // создание словаря где ключ - слово из среза, а значение - заглушка

	for _, word := range input { // если придет слово которое уже встречалось, то в словаре просто обновится значени
		m[word] = 1
	}
	for key := range m { // вывод всех ключей словаря
		fmt.Print(key, " ")
	}
}
