package problem1

import "fmt"

/*
 * 1. Введение в Go и установка - Задание 2
 * Настройте модуль с помощью команды go mod init myapp.
 * Добавьте функцию greet(name string) для вывода приветствия с именем.
 * Вызовите эту функцию из main().
 */

func Greet(name string) {
	fmt.Println("Привет", name)
}
