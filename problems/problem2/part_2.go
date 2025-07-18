package problem2

import "fmt"

/*
 * 2. Основы синтаксиса - Задание 2
 * Напишите программу с использованием switch, которая принимает оценку (1–5).
 * Выведите текстовое описание ("отлично", "хорошо" и т.д.).
 * Обработайте случай неверного ввода.
 */

func CheckNote() {
	var note int
	fmt.Println("Введите оценку от 1 до 5: ")
	fmt.Scanln(&note)
	switch note {
	case 1:
		fmt.Println("очень плохо")
	case 2:
		fmt.Println("плохо")
	case 3:
		fmt.Println("приемлемо")
	case 4:
		fmt.Println("хорошо")
	case 5:
		fmt.Println("очень хорошо")
	default:
		fmt.Println("Ошибка: Вы ввели некоректную оценку")
	}

}
