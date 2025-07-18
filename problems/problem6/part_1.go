package problem6

import "fmt"

/*
 * 6. Работа со срезами и массивами - Задание 1
 * Создайте срез из 5 чисел.
 * Добавьте в него 2 новых элемента с помощью append.
 * Выведите результат.
 */

func TestSliceAppend() {
	intSlice := []int{1, 2, 3, 4, 5}
	intSlice = append(intSlice, 1, 2)
	fmt.Println(intSlice)
}
