package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(numbers)) // добавление в WaitGroup горутин
	for _, num := range numbers {
		go square(num, &wg) // вызов горутин для каждого элемента списка
	}
	wg.Wait() // боликровка основного потока до окончания рабоыт всех горутин
}

// функция считающая квадрат числа
func square(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(number * number)
}
