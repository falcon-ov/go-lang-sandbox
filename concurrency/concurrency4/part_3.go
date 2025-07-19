package concurrency4

import (
	"fmt"
	"time"
)

func OneThread() {
	start := time.Now()
	total := 0
	const N = 1_000_000

	for i := 0; i < N; i++ {
		total += DoWork(i)
	}

	fmt.Println("Сумма:", total)
	fmt.Println("Время выполнения (последовательно):", time.Since(start))
}
