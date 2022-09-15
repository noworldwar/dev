package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	num    int64
	wg     sync.WaitGroup
	rwlock sync.RWMutex
	lock   sync.Mutex
)

func main() {
	start := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go Read()
	}

	wg.Wait()
	end := time.Now()

	fmt.Println(end.Sub(start))
}

func Write() {
	// rwlock.Lock()
	lock.Lock()
	num++
	time.Sleep(10 * time.Millisecond)

	lock.Unlock()
	// rwlock.Unlock()

	wg.Done()
}

func Read() {
	lock.Lock()
	// rwlock.RLock()

	time.Sleep(time.Millisecond)

	// rwlock.RUnlock()
	lock.Unlock()
	wg.Done()
}
