package problem11

import (
	"fmt"
	"time"
)

/*
 * 11. Параллелизм (Concurrency) - Задание 2
 * Создайте канал для передачи строк.
 * Запустите горутину, которая отправляет 5 слов в канал.
 * Выведите их в main().
 */

func TestChan() chan string {
	ch := make(chan string, 5)
	go func() {
		for i := range 5 {
			ch <- "word" + string(i)
		}
	}()
	time.Sleep(time.Second * 1)
	close(ch)
	return ch
}

func PrintChan(ch chan string) {
	for v := range ch {
		fmt.Println(v)
	}
}
