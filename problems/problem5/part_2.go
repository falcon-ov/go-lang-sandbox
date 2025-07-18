package problem5

/*
 * 5. Управление пакетами - Задание 2
 * Создайте пакет utils с функцией проверки чётности числа.
 * Используйте эту функцию в main для проверки введённого числа.
 */
//SUCCESS

func CheckEven(num int) (isEven bool) {
	if num%2 == 0 {
		isEven = true
		return
	}
	isEven = false
	return
}
