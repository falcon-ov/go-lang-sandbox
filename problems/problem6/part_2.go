package problem6

/*
 * 6. Работа со срезами и массивами - Задание 2
 * Напишите функцию, которая принимает срез чисел.
 * Возвращайте новый срез, содержащий только чётные числа.
 */

func ConvertToEven(slice []int) (sliceOnlyEven []int) {
	for _, v := range slice {
		if v%2 == 0 {
			sliceOnlyEven = append(sliceOnlyEven, v)
		}
	}
	return
}
