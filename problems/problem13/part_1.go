package problem13

import (
	"fmt"
	"time"
)

/*
 * 13. Стандартная библиотека - Задание 1
 * Используя пакет time, выведите текущую дату и время.
 */

func GetTime() {
	year, mounth, day := time.Now().Date()
	hours, minutes, seconds := time.Now().Clock()
	fmt.Println(year, mounth, day, hours, "H", minutes, "M", seconds, "S")
}
