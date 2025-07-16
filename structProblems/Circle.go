package structproblems

import "math"

type Circle struct {
	X, Y, R float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}
