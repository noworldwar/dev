package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	num  int64
	lock sync.Mutex
	wg   sync.WaitGroup
)

func add() {
	num++
	wg.Done()
}

func mutexAdd() {
	lock.Lock()
	num++
	lock.Unlock()
	wg.Done()
}

func atomicAdd() {
	atomic.AddInt64(&num, 1)
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 200000; i++ {
		wg.Add(1)
		go add()
		// go mutexAdd()
		// go atomicAdd()
	}

	wg.Wait()

	end := time.Now()

	fmt.Println(num)

	fmt.Println(end.Sub(start))
}
