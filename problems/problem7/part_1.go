package problem7

import "fmt"

/*
 * 7. Словари (map) - Задание 1
 * Создайте словарь, где ключами будут имена, а значениями — возраст.
 * Добавьте 3 записи и выведите их.
 */

func TestMap() {
	var simpleMap = make(map[string]int)
	simpleMap["Daniil"] = 20
	simpleMap["Oxana"] = 20
	simpleMap["Igori"] = 45
	fmt.Println(simpleMap)
}
