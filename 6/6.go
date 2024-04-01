package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	exit := make(chan int)
	go func() { // остановка горутины через канал
		defer wg.Done()
		for {
			select {
			case <-exit:
				return
			default:
				time.Sleep(10 * time.Millisecond)
				fmt.Println("Goroutine #1")
			}
		}
		fmt.Println()
	}()

	ctx, cancel := context.WithCancel(context.Background())

	go func() { // остановка горутины через context
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(10 * time.Millisecond)
				fmt.Println("Goroutine #2")
			}

		}
	}()

	ctx2, cancel2 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel2()

	go func() { // остановка горутины после таймаута
		defer wg.Done()
		for {
			select {
			case <-ctx2.Done():
				return
			default:
				time.Sleep(10 * time.Millisecond)
				fmt.Println("Goroutine #3")
			}
		}
	}()

	ch := make(chan int)

	go func() { // остановка горутины после закрытия канала
		defer wg.Done()
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					return
				}
			default:
				time.Sleep(10 * time.Millisecond)
				fmt.Println("Goroutine #4")
			}
		}
	}()

	p := new(bool)

	go func(p *bool) { // остановка горутины после изменения указателя
		defer wg.Done()
		for *p != true {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("Goroutine #5")
		}
		return
	}(p)

	time.Sleep(100 * time.Millisecond)
	*p = true
	exit <- 1
	cancel()
	close(ch)
	wg.Wait()
}
