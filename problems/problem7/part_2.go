package problem7

/*
 * 7. Словари (map) - Задание 2
 * Напишите функцию, которая принимает словарь (ключ — строка, значение — число).
 * Функция должна возвращать ключ с максимальным значением.
 */

func KeyMaxValue(simpleMap map[string]int) string {
	var pairIndexValue struct {
		Key   string
		Value int
	}
	for i, v := range simpleMap {
		if v > pairIndexValue.Value {
			pairIndexValue.Key = i
			pairIndexValue.Value = v
		}
	}
	return pairIndexValue.Key
}
