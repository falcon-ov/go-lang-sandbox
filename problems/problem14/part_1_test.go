package problem14

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		a, b, result int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{9, 5, 14},
	}
	for _, c := range cases {
		realResult := Sum(c.a, c.b)
		if realResult != c.result {
			t.Errorf("%v + %v ожидлось равно %v, фактически %v", c.a, c.b, c.result, realResult)
		}
	}
}
