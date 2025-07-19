package concurrency4

import (
	"fmt"
	"sync"
	"time"
)

func Parallel() {
	start := time.Now()
	const N = 1_000_000
	const workers = 4
	step := N / workers

	var wg sync.WaitGroup
	results := make([]int, workers)

	for w := 0; w < workers; w++ {
		from := w * step
		to := from + step

		wg.Add(1)
		go func(idx, from, to int) {
			defer wg.Done()
			sum := 0
			for i := from; i < to; i++ {
				sum += DoWork(i)
			}
			results[idx] = sum
		}(w, from, to)
	}

	wg.Wait()

	// собираем итог
	total := 0
	for _, val := range results {
		total += val
	}

	fmt.Println("Сумма:", total)
	fmt.Println("Время выполнения (горутины):", time.Since(start))
}
