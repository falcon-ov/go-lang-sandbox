package concurrency2

/*
Сложное задание:
Реализуйте программу, где две горутины отправляют числа в
один и тот же канал (например, одна отправляет четные числа, другая — нечетные).
Главная горутина должна принимать и суммировать все числа, пока канал не будет закрыт.
*/

func Conc2() int {
	ch := make(chan int)
	closeCh := make(chan bool)
	//Even
	go func(ch chan int, closeCh chan bool) {
		for i := range 10 {
			if i%2 == 0 {
				ch <- i
			}
		}
		closeCh <- true
	}(ch, closeCh)
	//Odd
	go func(ch chan int, closeCh chan bool) {
		for i := range 10 {
			if i%2 != 0 {
				ch <- i
			}
		}
		closeCh <- true
	}(ch, closeCh)

	go func(ch chan int, closeCh chan bool) {
		for {
			<-closeCh
			<-closeCh
			break
		}
		close(ch)
		close(closeCh)
	}(ch, closeCh)

	var sum int
	for v := range ch {
		sum += v
	}
	return sum
}
