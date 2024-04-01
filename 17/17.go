package main

import (
	"fmt"
)

func main() {
	// Инициализация отсортированного среза
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Вызов функции для поиска элемента 9 в срезе
	fmt.Println(binarySearch(arr, 9))
}

func binarySearch(arr []int, target int) int {
	left := 0             // Левая граница диапазона
	right := len(arr) - 1 // Правая граница диапазона
	var middle int        // Индекс среднего элемента
	// Пока левая граница не превышает правую
	for left <= right {
		middle = (left + right) / 2 // Вычисляем индекс среднего элемента
		// Если средний элемент меньше целевого
		if arr[middle] < target {
			left = middle + 1 // Изменяем левую границу диапазона
			// Если средний элемент больше целевого
		} else if arr[middle] > target {
			right = middle - 1 // Изменяем правую границу диапазона
			// Если средний элемент равен целевому
		} else if arr[middle] == target {
			return middle // Возвращаем индекс среднего элемента
		}
	}
	// Если целевой элемент не найден, возвращаем -1
	return -1
}
