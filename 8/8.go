package main

import "fmt"

func main() {
	var num int64
	var bit int
	fmt.Println("Введите число")
	fmt.Scan(&num)
	fmt.Println("В какой бит вы хотите установить 1?")
	fmt.Scan(&bit)
	var mask int64
	mask = 1 << bit                    // создание маски
	fmt.Println("Результат", num|mask) // применение операции поразрядной дизъюнкции
}
