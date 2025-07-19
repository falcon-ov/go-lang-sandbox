package concurrency4

import (
	"fmt"
	"sync"
	"time"
)

/*
Простое задание:
Напишите программу, где горутина отправляет в буферизированный канал 10 чисел,
а главная горутина читает их и выводит.
Убедитесь, что размер буфера не вызывает дедлок.
*/

func Conc1() {
	start := time.Now() // засекаем время
	ch := make(chan int, 10000000)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ch chan int, wg *sync.WaitGroup) {
		for i := range 10000000 {
			ch <- i
			// fmt.Println("inner go func:", i)
		}
		wg.Done()
		close(ch)
	}(ch, &wg)

	// for i := range ch {
	// 	fmt.Println(i)
	// }
	count := 0
	for i := range ch {
		count += i
	}
	fmt.Println("Сумма:", count)
	wg.Wait()
	duration := time.Since(start) // сколько прошло времени
	fmt.Println("Время выполнения:", duration)
}
