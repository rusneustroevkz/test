package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

const count = 100

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		ok()
		wg.Done()
	}()
	go func() {
		wrong()
		wg.Done()
	}()
	wg.Wait()
}

func ok() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	for i := 0; i < count; i++ {
		go doOk(ctx)
	}

	time.Sleep(time.Second * 1)

	cancel()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("ok", m.Alloc)

	time.Sleep(time.Second * 2)
}

func wrong() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	for i := 0; i < count; i++ {
		go do(ctx)
	}

	time.Sleep(time.Second * 1)

	cancel()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("wrong", m.Alloc)

	time.Sleep(time.Second * 2)
}

func do(ctx context.Context) {
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
	case <-ctx.Done():
		//fmt.Println("done")
	}
}

func doOk(ctx context.Context) {
	c := time.NewTimer(time.Second * 2)
	select {
	case <-c.C:
		fmt.Println("timeout")
	case <-ctx.Done():
		c.Stop()
		//fmt.Println("done")
	}
}
