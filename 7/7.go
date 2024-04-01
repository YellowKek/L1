package main

import (
	"fmt"
	"sync"
)

type SafetyMap struct { // создание структуры для безопасной записи в map
	sync.Mutex
	m map[int]int
}

func NewSafetyMap() *SafetyMap { // конструктор SafetyMap
	return &SafetyMap{
		m: make(map[int]int),
	}
}

func (m *SafetyMap) Load(key int, value int) { // метод для безопасного добавления
	m.Lock()
	defer m.Unlock()
	v, ok := m.m[key]
	if ok {
		m.m[key] = v + 1
	} else {
		m.m[key] = value
	}
}

func main() {
	sm := NewSafetyMap()
	wg := sync.WaitGroup{}
	wg.Add(4)
	go foo(sm, &wg)
	go foo(sm, &wg)
	go foo(sm, &wg)
	go func() { // функция для добавления 5 цифр в map
		defer wg.Done()
		for i := 0; i < 5; i++ {
			sm.Load(i, 1)
		}
	}()

	wg.Wait()
	fmt.Println(sm.m)
}

func foo(sm *SafetyMap, wg *sync.WaitGroup) { // функция для добавления 10 цифр в map
	defer wg.Done()
	for i := 0; i < 10; i++ {
		sm.Load(i, 1)
	}
}
