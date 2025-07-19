package concurrency1

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Сложное задание:
Создайте программу, которая запускает 10 горутин,
каждая из которых выводит свой порядковый номер и
случайное число от 1 до 100. Убедитесь, что вывод происходит
в порядке завершения горутин (потребуется изучить синхронизацию).
*/
func StartTenGoroutines() {
	for i := range 10 {
		go func(i int) {
			time.Sleep(time.Millisecond * time.Duration(i*rand.Intn(10)))
			fmt.Println(i, rand.Intn(100))
		}(i)
	}
}
