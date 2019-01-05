package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	var wg sync.WaitGroup
	m := make(chan string)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go compare(ctx, &wg, m, i)
	}

	wg.Wait()
	fmt.Println("close m")
	close(m)

	select {
	case m, ok := <-m:
		cancel()
		if !ok {
			fmt.Println("not found")
		}
		fmt.Println(m)
	}

}

func compare(ctx context.Context, wg *sync.WaitGroup, m chan<- string, i int) {
	defer wg.Done()
	time.Sleep(5 * time.Millisecond)
	fmt.Println(i)
	if i < 10 {
		m <- "wao!"
	}
}
