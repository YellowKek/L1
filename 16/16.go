package main

import "fmt"

func main() {
	// Инициализация среза
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	// Вызов функции для сортировки среза
	quickSort(0, len(arr)-1, arr)
	fmt.Println(arr)
}

func quickSort(left, right int, arr []int) {
	// Базовый случай: если длина среза равна 0, возвращаемся
	if len(arr) == 0 {
		return
	}
	// Если левая граница больше правой, возвращаемся
	if left > right {
		return
	}

	var (
		l     = left  // Индекс левой границы
		r     = right // Индекс правой границы
		pivot int     // Опорный элемент
	)
	// Выбираем опорный элемент как средний элемент
	pivot = (left + right) / 2
	// Пока левая граница меньше или равна правой
	for l <= r {
		// Поиск элемента слева, который больше опорного
		for arr[l] < arr[pivot] {
			l++
		}
		// Поиск элемента справа, который меньше опорного
		for arr[r] > arr[pivot] {
			r--
		}
		// Если индекс левой границы меньше или равен индексу правой границы
		if l <= r {
			// Меняем местами элементы на позициях l и r
			arr[l], arr[r] = arr[r], arr[l]
			// Сдвигаем границы
			l++
			r--
		}
		// Рекурсивный вызов quickSort для левой части среза
		if l < right {
			quickSort(left, r, arr)
		}
		// Рекурсивный вызов quickSort для правой части среза
		if r > left {
			quickSort(l, right, arr)
		}
	}
}
