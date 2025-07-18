package problem4

/*
 * 4. Функции - Задание 2
 * Напишите функцию minMax(a, b, c int) (int, int).
 * Функция должна возвращать минимальное и максимальное значение из трёх чисел.
 */

func MinMax(a ...int) (minValue int, maxValue int) {
	for _, v := range a {
		if v < minValue {
			minValue = v
		}
		if v > maxValue {
			maxValue = v
		}
	}
	return minValue, maxValue
}
