package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Counter struct { // описание структуры счетчика
	counter int32
}

func (c *Counter) Count(v int32) { // метод для увеличения счетчика
	atomic.AddInt32(&c.counter, v)
}

func main() {
	counter := Counter{}
	// запуска двух горутин
	go f(&counter)
	go f(&counter)
	time.Sleep(100 * time.Millisecond)
	fmt.Println(counter.counter)
}

func f(counter *Counter) { // каждая горутина увеличивает счетчик 10 раз на 10
	for i := 0; i < 10; i++ {
		counter.Count(10)
	}
}
