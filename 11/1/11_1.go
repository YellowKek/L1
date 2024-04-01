package main

import "fmt"

func main() {
	// создание множеств
	var set1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var set2 = []int{2, 3, 4, 5}

	m := make(map[int]int)

	for _, value := range set1 { // заполнение map значениями из set1,
		// так как словарь еще пусстой то все элементы из set1 еще ни разу не встречались,
		//добавляем значение в map и ставим счетчик на 1
		m[value] = 1
	}

	for _, value := range set2 { // если значение уже есть в словаре, увеличиваем счетчик на 1
		v, ok := m[value]
		if !ok {
			m[value] = 1
		} else {
			m[value] = v + 1
		}
	}

	for key, value := range m {
		if value == 2 { // если значение встретилось ровно два раза, значит оно есть в обоих мнржествах
			fmt.Print(key, " ")
		}
	}
}
