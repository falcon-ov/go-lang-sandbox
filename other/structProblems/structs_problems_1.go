package structproblems

import (
	"math"
)

func OldCircleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func CircleArea(c Circle) float64 {
	return math.Pi * c.R * c.R
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

/*
1. Какая разница между методом и функцией?
	I. Метод это функция, которая привязана к определенной структуре

2. В каких случаях могут пригодиться встроенные (скрытые) поля?
	I. Когда мы хотим, чтобы другая структура наследовала все полня и методы другой структуры

3. Добавьте новый метод perimeter в интерфейс Shape,
который будет вычислять периметр фигуры.
Имплементируйте этот метод для Circle и Rectangle.
	I. Реализовал
*/
