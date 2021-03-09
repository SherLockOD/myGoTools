package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context 通知所有 goroutine 关闭
func main() {
	wg := sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("work%d", i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(ctx, name)
		}()
	}
	<- time.After(time.Second * 10)
	cancelFunc()

	wg.Wait()
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("quit", name)
			return
		default:
			<- time.After(time.Second * 1)
		}
		fmt.Println(name, time.Now().String())
	}
}