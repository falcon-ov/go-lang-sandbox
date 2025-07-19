package concurrency2

import "fmt"

/*
Простое задание:
Напишите программу, где одна горутина отправляет в канал 5 строк
(например, "Сообщение 1", "Сообщение 2" и т.д.),
а главная горутина их принимает и выводит.
Закройте канал после отправки всех сообщений.
*/
func Conc1() {
	ch := make(chan string)
	go func(ch chan string) {
		for v := range 5 {
			ch <- fmt.Sprintf("Сообщение %v", v)
		}
		close(ch)
	}(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
