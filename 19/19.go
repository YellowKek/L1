package main

import "fmt"

func main() {
	var word = "главрыба"
	temp := []rune(word) // создание слайса рун из строки

	for i := 0; i < len(temp)/2; i++ { // проход до длины строки / 2
		temp[i], temp[len(temp)-i-1] = temp[len(temp)-i-1], temp[i] // меняю символы местами
	}
	fmt.Println(string(temp))
}
