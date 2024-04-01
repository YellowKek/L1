package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// создание каналов
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {
		for _, number := range numbers { // запись в канал данных из списка
			ch1 <- number
		}
		close(ch1) // после того как все данные были отправлены в канал, закрываем этот канал
	}()

	go func() {
		for i := range ch1 { // проход по всем данным из канала и передача их в следующий канал
			ch2 <- i * 2
		}
		close(ch2) // после того как все данные были отправлены в канал, закрываем этот канал
	}()

	for i := range ch2 { // вывод данных из второго канала
		fmt.Println(i)
	}
}
