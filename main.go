package main

import (
	"fmt"
	"math"
	"sync"
)

const DealNumPerGoroutine = 1000

func main() {
	var syncMap sync.Map
	sum := 2000001
	for i := 1; i < sum; i++ {
		syncMap.Store(i, i)
	}

	goroutineNum := math.Ceil(float64(sum / DealNumPerGoroutine))
	var wg sync.WaitGroup
	for i := 1; i <= int(goroutineNum); i++ {
		wg.Add(1)
		go func(i int, pool *sync.Map) {
			defer wg.Done()
			myGoroutineNew(i, &syncMap)
		}(i, &syncMap)
	}
	wg.Wait()

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:", key, value)
		return true
	})

}

func myGoroutineNew(i int, pool *sync.Map) {
	var start int
	start = (i-1)*DealNumPerGoroutine + 1
	end := start + DealNumPerGoroutine
	for j := start; j < end; j++ {
		fmt.Printf("%v\t\t", j)
		pool.Delete(j)
	}
	fmt.Println()
}
