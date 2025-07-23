package ex2

import (
	"context"
	"sync"
)

//реализовать функцию, которая собирает данные из массива каналов
//в один канал, предусмотреть возможость отмены контекста

func FanIn(ctx context.Context, chans []chan int) chan int {
	out := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}

		for _, ch := range chans {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case v, ok := <-ch:
						if !ok {
							return
						}
						select {
						case out <- v:
						case <-ctx.Done():
							return
						}
					case <-ctx.Done():
						return
					}
				}
			}()
			wg.Wait()
			close(out)
		}
	}()
	return out
}
