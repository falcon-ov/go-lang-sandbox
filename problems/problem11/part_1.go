package problem11

import (
	"fmt"
	"time"
)

/*
 * 11. Параллелизм (Concurrency) - Задание 1
 * Напишите программу с двумя горутинами.
 * Одна выводит чётные числа от 1 до 10, другая — нечётные.
 * Используйте time.Sleep для синхронизации.
 */

func PrintOdds() {
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
			time.Sleep(time.Second*1 + time.Microsecond*500)
		}
	}
}
func PrintEvens() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			time.Sleep(time.Second * 1)
		}
	}
}

func TestGor() {
	go PrintEvens()
	go PrintOdds()
}
