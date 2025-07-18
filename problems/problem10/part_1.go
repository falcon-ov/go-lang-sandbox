package problem10

import "errors"

/*
 * 10. Обработка ошибок - Задание 1
 * Напишите функцию divide(a, b int) (int, error).
 * Функция должна возвращать ошибку при делении на 0.
 */

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

/*
	val, err := problem10.Divide(6, 2)
	fmt.Println("Value = ", val)
	fmt.Println("Err =", err)
*/
