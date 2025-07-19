package concurrency3

import (
	"sync"
)

/*
Сложное задание:
Напишите программу, которая запускает N горутин (N задается пользователем),
каждая из которых выполняет случайную задачу (например, вычисляет сумму чисел от 1 до 1000).
Используйте WaitGroup для ожидания завершения всех горутин и соберите результаты в главном потоке.
*/
type WokersJob struct {
	WorkerId int
	Result   int
}

func Conc2(n int) chan WokersJob {
	var wg sync.WaitGroup
	ch := make(chan WokersJob, n)

	wg.Add(n)
	for i := range n {
		go func(i int, ch chan WokersJob, wg *sync.WaitGroup) {
			totalsum := 0
			for j := range 1000 {
				totalsum += j
			}
			ch <- WokersJob{i, totalsum}
			wg.Done()
		}(i, ch, &wg)
	}
	wg.Wait()
	close(ch)
	return ch
}
