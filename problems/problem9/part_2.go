package problem9

import (
	"fmt"
	"math"
)

/*
 * 9. Интерфейсы - Задание 2
 * Создайте интерфейс Shape с методом Area() float64.
 * Реализуйте его для структур Circle и Rectangle.
 * Вычислите площади фигур.
 */
type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}
type Rectangle struct {
	Lenght float64
	Width  float64
}

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}
func (r Rectangle) Area() float64 {
	return r.Lenght * r.Width
}

func PrintArea(shape Shape) {
	fmt.Println("Площадь =", shape.Area())
}

func TestInterface2() {
	var c1 = Circle{10}
	var r1 = Rectangle{5, 5}

	PrintArea(c1)
	PrintArea(r1)
}
