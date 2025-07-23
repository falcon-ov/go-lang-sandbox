package final2

import (
	"context"
	"fmt"
	"time"
)

/*
Создайте программу, которая имитирует выполнение задачи (например, чтение данных) в
горутине с длительностью 7 секунд. Используйте context.WithTimeout,
чтобы прервать задачу через 4 секунды, и выведите сообщение об ошибке таймаута.
*/

func Ex1() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	go func() {
		fmt.Println("Работа горутины начата")
		time.Sleep(time.Second * 7)
		fmt.Println("Работа горутины завершена")
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Задача отмена по таймауту", ctx.Err())
	}
}
