package concurrency1

import (
	"fmt"
	"time"
)

func PrintOdds() {
	time.Sleep(time.Millisecond * 10)
	for i := range 100 {
		if i%2 != 0 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 10)
		}
	}
}
func PrintEvens() {
	time.Sleep(time.Millisecond * 5)
	for i := range 100 {
		if i%2 == 0 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func TestGor() {
	go PrintEvens()
	go PrintOdds()
}
