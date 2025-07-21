package concurrency5

import (
	"fmt"
	"strconv"
)

/*
Простое задание:
Напишите программу, которая использует select для чтения из двух каналов:
один передает числа, другой — строки.
Выводите сообщения по мере их поступления.
*/

func GetWords(n int, chInt chan int) {
	for i := range n {
		chInt <- i
	}
	close(chInt)
}
func GetNums(n int, chString chan string) {
	for i := range n {
		chString <- "word" + strconv.Itoa(i)
	}
	close(chString)
}

func Conc1() {
	chInt := make(chan int)
	chString := make(chan string)
	go GetNums(10, chString)
	go GetWords(10, chInt)
	closedCount := 0

	for {
		select {
		case msg1, ok := <-chInt:
			if ok {
				fmt.Println("Получено из chInt: ", msg1)
			} else {
				closedCount++
				chInt = nil
			}

		case msg2, ok := <-chString:
			if ok {
				fmt.Println("Получено из chString: ", msg2)
			} else {
				closedCount++
				chString = nil
			}

		}
		if closedCount == 2 {
			break
		}
	}
}
