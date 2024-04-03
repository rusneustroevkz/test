package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const mapSliceSpeed = 1_000_000

var slice []int
var mapa = make(map[int]struct{})

func init() {
	for i := 0; i < mapSliceSpeed; i++ {
		slice = append(slice, i)
		mapa[i] = struct{}{}
	}
}

func main() {
	cur := rand.Intn(mapSliceSpeed)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		s := time.Now()
		for _, vv := range slice {
			if cur == vv {
				break
			}
		}
		fmt.Printf(" %d", time.Since(s).Nanoseconds())
	}()

	go func() {
		defer wg.Done()
		s := time.Now()
		_ = mapa[cur]
		fmt.Printf("map %d, ", time.Since(s).Nanoseconds())
	}()

	wg.Wait()
}
