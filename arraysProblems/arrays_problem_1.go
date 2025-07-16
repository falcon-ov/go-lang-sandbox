package arraysProblems

import "fmt"

// Напишите программу, которая находит самый наименьший элемент в этом списке:

func MinElement() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	var min = x[0]
	for i := 1; i < len(x); i++ {
		if min > x[i] {
			min = x[i]
		}
	}

	fmt.Println(min)

}
