package concurrency3

import (
	"fmt"
	"sync"
	"time"
)

/*
Простое задание: Модифицируйте пример из темы 1 (с четными и нечетными числами),
 используя sync.WaitGroup вместо time.Sleep для ожидания завершения горутин.
*/

func PrintOdds(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 10)
	for i := range 100 {
		if i%2 != 0 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 10)
		}
	}
}
func PrintEvens(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 5)
	for i := range 100 {
		if i%2 == 0 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func TestGor() {
	var wg sync.WaitGroup
	wg.Add(2)
	go PrintEvens(&wg)
	go PrintOdds(&wg)
	wg.Wait()
}
