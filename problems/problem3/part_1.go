package problem3

import "fmt"

/*
 * 3. Типы данных - Задание 1
 * Создайте переменные типов int, float64, string, bool.
 * Выведите их значения и типы, используя fmt.Printf с %T.
 */

func CreateSomeTypes() {
	var i int = 5
	var f float64 = 2.15
	var s string = "StRiNg"
	var b bool = true
	values := []any{i, f, s, b}
	for _, variable := range values {
		fmt.Printf("%T = %v \n", variable, variable)
	}

}
