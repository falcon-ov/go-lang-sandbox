package concurrency4

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Сложное задание:
Реализуйте "конвейер" из трех горутин: первая генерирует числа, вторая удваивает их,
третья выводит результат.
Используйте буферизированные каналы для передачи данных между горутинами.
*/

func Conver() {
	n := 100 //кол-во генерчисел
	chGen := make(chan int, n)
	chDouble := make(chan int, n)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(n int, chGen chan int, wg *sync.WaitGroup) {
		for i := range n {
			chGen <- rand.Intn(100) + i
		}
		close(chGen)
	}(n, chGen, &wg)

	go func(chGen chan int, chDouble chan int, wg *sync.WaitGroup) {
		for v := range chGen {
			chDouble <- v * 2
		}
		close(chDouble)
	}(chGen, chDouble, &wg)

	go func(chDouble chan int, wg *sync.WaitGroup) {
		for v := range chDouble {
			fmt.Println(v)
		}
		wg.Done()
	}(chDouble, &wg)
	wg.Wait()
}
