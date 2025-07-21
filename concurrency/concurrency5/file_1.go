package concurrency5

import (
	"fmt"
	"sync"
	"time"
)

/*
Написать 3 функции: Функции должны работать одновременно
writer - генерит числа от 1 до 10
doubler - умножает числа на 2, имитируя работу (500 ms)
reader - читает и выводит на экран
*/

func writer() <-chan int {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ch chan int) {
		for i := range 10 {
			ch <- i + 1
		}
		close(ch)
	}(ch)
	return ch
}
func doubler(ch <-chan int) <-chan int {
	chDouble := make(chan int)
	go func(ch <-chan int) {
		for v := range ch {
			time.Sleep(time.Millisecond * 500)
			chDouble <- v * 2
		}
		close(chDouble)
	}(ch)
	return chDouble
}
func reader(chDouble <-chan int) {

	for v := range chDouble {
		fmt.Println(v)
	}

}

func Test() {
	reader(doubler(writer()))
}
