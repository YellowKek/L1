package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

func main() {
	fmt.Println("Введите кол-во ворекров")
	var n int
	fmt.Scan(&n)

	wg := sync.WaitGroup{}
	wg.Add(n)
	ch := make(chan int, 100) // создание буферизированного канала который принимает int
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < n; i++ {
		go printNum(ch, ctx, &wg) // запуск горутин
	}

	c := make(chan os.Signal) // создание канала который будет ожидать сигнал Ctrl + C
	signal.Notify(c, os.Interrupt)
	go func() {
		for {
			select {
			case <-ctx.Done(): // ожидание вызова cancel функции
				return
			default:
				rnd := rand.Intn(100) // генерация случайных чисел
				ch <- rnd             // отправка чисел в канал
			}
		}
	}()
	<-c       // ожидание сигнала
	cancel()  // вызов cancel функции
	wg.Wait() // ожидание завершения всех горутин
}

func printNum(ch chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // ожидание вызова cancel функции
			fmt.Println("worker done")
			return
		case data := <-ch: // ожидание чисел из канала
			fmt.Println(data)
		}
	}
}
