package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	r := 1_000_000

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		c := 0
		mu := sync.Mutex{}
		start := time.Now()
		for i := 0; i < r; i++ {
			func() {
				mu.Lock()
				defer mu.Unlock()
				c++
			}()
		}
		fmt.Println("mu", time.Since(start))
	}()

	go func() {
		defer wg.Done()
		c := 0
		ch := make(chan int)
		start := time.Now()
		go func() {
			for i := 0; i < r; i++ {
				ch <- i
			}
			close(ch)
		}()
		for {
			select {
			case x, ok := <-ch:
				if !ok {
					fmt.Println("ch", time.Since(start))
					return
				}
				c += x
			}
		}

	}()

	go func() {
		defer wg.Done()
		var c atomic.Int64
		start := time.Now()
		for i := 0; i < r; i++ {
			func() {
				c.Add(1)
			}()
		}
		fmt.Println("at", time.Since(start))
	}()

	wg.Wait()
}
