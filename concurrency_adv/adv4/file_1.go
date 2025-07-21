package adv4

import (
	"fmt"
	"sync"
)

/*
4. Паттерны конкурентности
Вопрос: Как работает паттерн "select" с каналами, и для чего он нужен?
Паттерн select нужен для неблокируещего приему данных из нескольких каналов,
работает он следующим образом есть два case <-ch1: case <-ch2: из них какой первый
канал будет готов, тот case и будет выбран для выполнения, если готовы оба выбирается случайный case,
если ни один не готов будет паника deadlock, если при этом есть default, то deadlock удасться
избежать и будет выполнен блок кода default

Задача: Напиши код с использованием select, где одна горутина ждёт данные из двух каналов и выводит первое полученное значение.
*/

func Ex4() *sync.WaitGroup {
	var wg sync.WaitGroup
	ch1, ch2 := make(chan int), make(chan int)
	wg.Add(1)
	go func(ch1 chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
		wg.Done()
	}(ch1, &wg)

	wg.Add(1)
	go func(ch2 chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 100; i++ {
			ch2 <- i
		}
		close(ch2)
		wg.Done()
	}(ch2, &wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup, ch1 chan int, ch2 chan int) {
		for {
			select {
			case msg1, ok1 := <-ch1:
				if ok1 {
					fmt.Println("ch1: ", msg1)
				} else {
					ch1 = nil
				}
			case msg2, ok2 := <-ch2:
				if ok2 {
					fmt.Println("ch2: ", msg2)
				} else {
					ch2 = nil
				}
			default:
				if ch1 == nil && ch2 == nil {
					wg.Done()
					return
				}
			}
		}
	}(&wg, ch1, ch2)

	return &wg
}
