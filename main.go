package main

import (
	"fmt"
	"time"
)

func main() {
	data := []string{"apple", "banana", "cherry"}

	for _, item := range data {
		go func() {
			time.Sleep(time.Second * 1)
			fmt.Println(item) //проблема в замыкании
		}()
	}

	time.Sleep(time.Second * 2) //плохая практика, лучше использвать waitgroups
}
