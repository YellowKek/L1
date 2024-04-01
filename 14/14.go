package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} // создание переменной любого типа
	a = 1
	//a = "string"
	//a = 1.1
	f(a)
}

func f(a any) {
	fmt.Println(reflect.TypeOf(a)) // определение типа
}
