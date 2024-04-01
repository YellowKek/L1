package main

import (
	"fmt"
	"math"
)

type Point struct { // описание структуры Point
	x int
	y int
}

func NewPoint(x int, y int) *Point { // конструктор Point
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) GetDistance(anotherPoint *Point) float64 { // метод который считает расстояние от точки на который был вызван метод до переданной точки
	return math.Sqrt(math.Pow(float64(anotherPoint.x-p.x), 2) + math.Pow(float64(anotherPoint.y-p.y), 2))
}

func main() {
	// создание двух точек
	point1 := NewPoint(1, 1)
	point2 := NewPoint(3, 3)
	fmt.Println(point1.GetDistance(point2)) // вывод расстояния от point1 до point2
}
