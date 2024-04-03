package main

import (
	"fmt"
	"slices"
	"sync"
	"time"
)

const maxRange = 1_0

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	a := []int{}
	m := make(map[int]struct{})

	ach := make(chan []int)
	mch := make(chan map[int]struct{})

	go func() {
		defer wg.Done()

		for i := 0; i < maxRange; i++ {
			a = append(a, i)
		}
		ach <- a
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < maxRange; i++ {
			m[i] = struct{}{}
		}
		mch <- m
	}()

	wg.Add(2)
	aa := <-ach
	go func() {
		defer wg.Done()
		start := time.Now()

		for i := range aa {
			if i == len(aa) {
				break
			}
			aa = slices.Delete(aa, i, i+1)
		}

		fmt.Println("slice", time.Since(start).Nanoseconds())
	}()

	mm := <-mch
	go func() {
		defer wg.Done()
		start := time.Now()

		for key := range mm {
			delete(mm, key)
		}

		fmt.Println("map", time.Since(start).Nanoseconds())
	}()
	wg.Wait()
}
