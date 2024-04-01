package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "snow dog sun"
	splited := strings.Split(text, " ")      // разделение строки по пробелам
	for i := len(splited) - 1; i >= 0; i-- { // вывод слов в обратном порядке
		fmt.Print(splited[i], " ")
	}
}
