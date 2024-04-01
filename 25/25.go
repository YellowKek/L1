package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	var sec time.Duration
	fmt.Println("Введите кол-во секунд")
	fmt.Scan(&sec)
	go func() {
		for {
			select {
			case <-time.After(sec * time.Second):
				c <- 1
			}
		}
	}()

	<-c
	fmt.Println("The end")
}
