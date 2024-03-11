package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(numbers)) // добавление в WaitGroup горутин
	var sum = 0
	for _, number := range numbers {
		go square(number, &sum, &wg) // вызов горутин для каждого элемента списка
	}
	wg.Wait() // боликровка основного потока до окончания рабоыт всех горутин
	fmt.Println(sum)
}

// функция считающая квадрат числа
func square(number int, sum *int, wg *sync.WaitGroup) {
	*sum += number * number // подсчет суммы квадратов чисел
	wg.Done()
}
