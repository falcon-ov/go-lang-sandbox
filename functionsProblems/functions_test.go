package functions

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

type testpairMax struct {
	values  []int
	average int
}

var testsMax = []testpairMax{
	{[]int{1, 2}, 2},
	{[]int{1, 3, 1, 4, 7, 1}, 7},
	{[]int{-1, 9}, 9},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMaxNum(t *testing.T) {
	for _, pair := range testsMax {
		v := MaxNum(pair.values...)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
