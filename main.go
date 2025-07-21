package main

import (
	"fmt"
	"math/rand"
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
	start := time.Now()

	for range 10 {
		totalWorkSeconds = +randomWait()
	}
	mainSeconds = int(time.Since(start).Seconds())
	fmt.Println("main: ", mainSeconds, "sec")
	fmt.Println("total: ", totalWorkSeconds, "sec")
}
