package structproblems

type Rectangle struct {
	X1, Y1, X2, Y2 float64
}

func (r *Rectangle) Area() float64 {
	l := distance(r.X1, r.Y1, r.X1, r.Y2)
	w := distance(r.X1, r.Y1, r.X2, r.Y1)
	return l * w
}

func (r *Rectangle) Perimeter() float64 {
	l := distance(r.X1, r.Y1, r.X1, r.Y2)
	w := distance(r.X1, r.Y1, r.X2, r.Y1)
	return 2*w + 2*l
}
