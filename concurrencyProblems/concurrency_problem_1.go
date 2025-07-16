package concurrencyProblems

import (
	"fmt"
	"math/rand"
	"time"
)

func F(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
func Ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func Pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func Printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

/*
1. Как задать направление канала?
например
var c chan<- string - значит создать канал только для отправки
var c <-chan string - значит создать канал только для принятия

2.Напишите собственную функцию Sleep, используя time.After
*/
func Sleep(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}

/*
3.Что такое буферизированный канал? Как создать такой канал с ёмкостью в 20 сообщений?
	буферизированный канал позволяет выполнять корутины асинхронно, то есть
*/
