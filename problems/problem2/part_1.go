package problem2

import "fmt"

/*
 * 2. Основы синтаксиса - Задание 1
 * Напишите программу, которая запрашивает у пользователя два числа.
 * Выведите их сумму, разность и произведение, используя fmt.Scanln.
 */

func SimpleCalc() {
	var x, y float64
	fmt.Println("Введите два числа(x, y) через пробел: ")
	fmt.Scanln(&x, &y)
	fmt.Println(x, " + ", y, " = ", x+y)
	fmt.Println(x, " - ", y, " = ", x-y)
	fmt.Println(x, " * ", y, " = ", x*y)
	if y == 0 {
		fmt.Println("y = 0, ошибка деления на ноль")
		return
	} else {
		fmt.Println(x, " / ", y, " = ", x/y)
	}

}
