package map_slice_write_speed

import (
	"fmt"
	"sync"
	"time"
)

const maxRange = 10

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var a []int
	m := make(map[int]struct{})

	go func() {
		defer wg.Done()

		start := time.Now()
		for i := 0; i < maxRange; i++ {
			a = append(a, i)
		}
		fmt.Println("array", time.Since(start).Nanoseconds())
	}()

	go func() {
		defer wg.Done()

		start := time.Now()
		for i := 0; i < maxRange; i++ {
			m[i] = struct{}{}
		}
		fmt.Println("map", time.Since(start).Nanoseconds())
	}()

	wg.Wait()
}
