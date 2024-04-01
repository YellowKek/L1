package main

import "fmt"

func main() {
	// объявление перменных
	var i int
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(list)
	fmt.Println("Введите инлекс элемента который хотите удалить")
	fmt.Scan(&i)                 // считываение индекса из консоли
	if i < 0 || i >= len(list) { // проверка индекса на валидность
		fmt.Println("Индекс вышел за границы массива")
	} else {
		list = append(list[:i], list[i+1:]...) // удаление элемента из списка
		fmt.Println(list)
	}
}
