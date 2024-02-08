package main

import (
	"fmt"
	"math"
)

// Структура Point с инкапсулированными параметрами x и y
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Метод для вычисления расстояния между двумя точками
func (p Point) DistanceTo(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func task24() {
	// Используем конструктор для создания двух точек
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	// Вычисляем расстояние между точками
	distance := point1.DistanceTo(point2)

	// Выводим результат
	fmt.Printf("Точка 1: (%.2f, %.2f)\n", point1.x, point1.y)
	fmt.Printf("Точка 2: (%.2f, %.2f)\n", point2.x, point2.y)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
