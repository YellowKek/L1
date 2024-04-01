package main

import (
	"fmt"
)

// Dog Определение структуры для собаки
type Dog struct {
}

// Cat Определение структуры для кошки
type Cat struct {
}

// AnimalAdapter Определение интерфейса с методом speak()
type AnimalAdapter interface {
	speak()
}

// DogAdapter Определение структуры, реализующей интерфейс AnimalAdapter
type DogAdapter struct {
	*Dog // Встраиваем структуру Dog
}

// Метод speak() для DogAdapter, который делегирует вызов метода speak() структуре Dog
func (d *Dog) speak() {
	fmt.Println("Wof")
}

// CatAdapter Определение структуры, реализующей интерфейс AnimalAdapter
type CatAdapter struct {
	*Cat // Встраиваем структуру Cat
}

// Метод speak() для CatAdapter, который делегирует вызов метода speak() структуре Cat
func (c *Cat) speak() {
	fmt.Println("Meow")
}

func main() {
	// Создание адаптера для собаки
	dogAdapter := DogAdapter{&Dog{}}
	// Создание адаптера для кошки
	catAdapter := CatAdapter{&Cat{}}
	// Вызов метода speak() для кошки
	catAdapter.speak()
	// Вызов метода speak() для собаки
	dogAdapter.speak()
}
