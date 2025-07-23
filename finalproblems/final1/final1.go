package final1

import (
	"context"
	"fmt"
	"time"
)

/*
Напишите программу, которая запускает горутину для вывода чисел от 1 до 10
с интервалом в 1 секунду. Используйте context.WithCancel,
чтобы прервать выполнение горутины через 5 секунд.
*/

func Ex1() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for i := range 10 {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена контекстом")
				return
			default:
				fmt.Println(i)
				time.Sleep(time.Second * 1)
			}
		}
	}()

	time.Sleep(time.Second * 5)
	cancel()
	time.Sleep(time.Millisecond * 500)
}
