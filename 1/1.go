package main

import "fmt"

type Human struct { // описание структуры Human
	Name string
	Age  uint8
}

func (h Human) Walk() string { // метод структуры Human который возварщает строку
	return "walk"
}

type Action struct { // описание структуры Action в которую встраивается структура Human
	Human
}

func NewAction(name string, age uint8) *Action { // конструктор Action
	return &Action{
		Human{
			Name: name,
			Age:  age,
		},
	}
}

func (a Action) doAction() { // метод структуры Action который выводит данные из встроенной структуры Human
	fmt.Println("human " + a.Name + " " + a.Walk())
}

func main() {
	action := NewAction("Ivan", 20) // создание объекта структуры Action
	action.doAction()
}
