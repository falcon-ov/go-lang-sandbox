package concurrency6

import (
	"fmt"
	"math/rand"
	"time"
)

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
}
func predictableTimeWork(fn func()) {
	timer := time.NewTimer(5 * time.Second)
	ch := make(chan string)
	go func(fn func()) {
		fn()
		ch <- "done"
	}(fn)
	select {
	case <-timer.C:
		fmt.Println("TimeOut")
	case <-ch:
		fmt.Println("Function done")
	}
}

func Test() {
	predictableTimeWork(randomTimeWork)
}
