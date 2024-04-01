package main

import "fmt"

func main() {
	// объявление переменных
	var (
		a = 5
		b = 10
	)

	a, b = b, a // обмен значениями
	fmt.Println(a, b)
}
