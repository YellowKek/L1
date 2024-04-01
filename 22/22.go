package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"log"
)

func main() {
	var ( // объявление переменных
		a, b   decimal.Decimal
		d1, d2 string
	)

	fmt.Println("Введите перове число")
	fmt.Scan(&d1) // считывание из консоли первого числа

	fmt.Println("Введите второе число")
	fmt.Scan(&d2) // считывание из консоли второго числа

	// преобразование чисел в decimal
	a, err := decimal.NewFromString(d1)
	b, err = decimal.NewFromString(d2)

	if err != nil {
		log.Fatal(err.Error())
	}

	// выполнение математических операций над числами
	fmt.Println("Результат умножения", a.Mul(b))
	fmt.Println("Результат деления", a.Div(b))
	fmt.Println("Результат сложения", a.Add(b))
	fmt.Println("Результат вычитания", a.Sub(b))

}
