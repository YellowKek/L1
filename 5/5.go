package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Введите кол-во секунд")
	var sec time.Duration
	fmt.Scan(&sec)
	ctx, cancel := context.WithCancel(context.Background())
	numbers := make([]int, 0)
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int, 100)
	go publish(ch, sec, cancel, &wg)
	go recieve(ch, ctx, &numbers, &wg)
	wg.Wait()
	//fmt.Println(numbers)
}

// функция которая отправляет данные в канал
func publish(ch chan int, sec time.Duration, cancelFunc context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.After(sec * time.Second)
	for {
		select {
		case <-t: // если время истекло вызываем cancel функцию
			cancelFunc()
			fmt.Println("time is up")
			return
		default: // если время еще не истекло, отправляем в канал случайные числа
			rnd := rand.Intn(100)
			ch <- rnd
		}
	}
}

// функция которая принимает данные из канала
func recieve(ch chan int, ctx context.Context, numbers *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return // если была вызвана cancel функция, горутина завершается
		default:
			*numbers = append(*numbers, <-ch) // пока не была вызвана cancel функция, принимаем данные из канала
		}
	}
}
