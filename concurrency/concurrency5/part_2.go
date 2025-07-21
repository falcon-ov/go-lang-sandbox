package concurrency5

import (
	"fmt"
	"time"
)

/*
2. **Сложное задание**:
Реализуйте программу, где горутина отправляет данные в два канала с разной скоростью,
 а главная горутина использует `select` с таймаутом (через `time.After`),
  чтобы завершить ожидание, если данные не поступают в течение 3 секунд.
*/

func Send2Chan(n int, ch1 chan int, ch2 chan int) {
	for i := range n {
		time.Sleep(time.Microsecond * 200)
		ch1 <- i
		time.Sleep(time.Microsecond * 100)
		ch2 <- i
	}
}

func GeneralFunction() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	breakBool := false

	go Send2Chan(10, ch1, ch2)

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(time.Second * 3):
			fmt.Println("Timeout: 3 seconds")
			breakBool = true
		}
		if breakBool {
			break
		}
	}
}
