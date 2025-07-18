package problem3

import "fmt"

/*
 * 3. Типы данных - Задание 2
 * Создайте массив из 5 чисел и срез из 3 строк.
 * Выведите их содержимое и измените один элемент в каждом.
 */

func CreateSomeArrays() {
	intArray := [5]int{1, 2, 3, 4, 5}
	stringSlice := []string{"hello", "world", "go"}
	fmt.Println(intArray)
	fmt.Println(stringSlice)

	intArray[2] = 0
	stringSlice[2] = "GO"

	fmt.Println(intArray)
	fmt.Println(stringSlice)
}
