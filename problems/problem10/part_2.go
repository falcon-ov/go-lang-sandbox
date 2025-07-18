package problem10

import "fmt"

/*
 * 10. Обработка ошибок - Задание 2
 * Напишите программу, которая запрашивает возраст.
 * Возвращайте ошибку, если введено отрицательное значение.
 * Обработайте ошибку в main().
 */

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Ошибка: %v", e.msg)
}

func ScanAge() error {
	fmt.Println("Введите возраст")
	var age int
	fmt.Scanln(&age)
	if age < 0 {
		return &MyError{"возраст меньше нуля"}
	}
	return nil
}

func TestErrors() {
	err := ScanAge()
	if err != nil {
		fmt.Println(err.Error())
	}
}
