package adv3

import (
	"fmt"
	"sync"
	"time"
)

/*
3. Синхронизация
Вопрос: Когда следует использовать sync.RWMutex вместо обычного sync.Mutex?
Ответ: Когда нам не важен порядок чтения данных из канала, и больше идет чтения, а не записи в
канал, то имеет смысл использовать sync.RWMutex, в случае, если у нас только запись
то sync.RWMutex может замедлить программу. sync.Mutex нужно исползовать, когда больше записей,
чем чтений


Задача: Напиши код, где несколько горутин читают общую переменную counter,
а одна горутина её увеличивает, используя sync.RWMutex.
*/

func Ex3() *sync.WaitGroup {
	var counter int
	var rwmu sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func(counter *int, rwmu *sync.RWMutex, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 5)
			rwmu.Lock()
			*counter++
			rwmu.Unlock()
		}
	}(&counter, &rwmu, &wg)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(counter *int, rwmu *sync.RWMutex, wg *sync.WaitGroup, i *int) {
			defer wg.Done()
			for {
				if *counter == 100 {
					break
				}
				rwmu.RLock()
				fmt.Println(*i, "-Goroutine: ", *counter)
				rwmu.RUnlock()
			}
		}(&counter, &rwmu, &wg, &i)
	}
	return &wg
}
