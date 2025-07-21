package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randomWait() int {
	workSeconds := rand.Intn(5 + 1)
	time.Sleep(time.Duration(workSeconds) * time.Second)
	return workSeconds
}

func main() {
	mainSeconds := 0
	totalWorkSeconds := 0
	wg := &sync.WaitGroup{}
	locker := &sync.Mutex{}
	start := time.Now()

	wg.Add(10)
	for range 10 {
		go func() {
			defer wg.Done()
			seconds := randomWait()
			locker.Lock()
			totalWorkSeconds += seconds
			locker.Unlock()
		}()
	}
	wg.Wait()
	mainSeconds = int(time.Since(start).Seconds())
	fmt.Println("main: ", mainSeconds, "sec")
	fmt.Println("total: ", totalWorkSeconds, "sec")
}
