package main

import (
	"fmt"
	"sync"
)

// func main() {

// 	data := []int{1, 2, 3, 5, 4, 9, 8, 4, 5}
// 	loopData := func(handleData chan<- int) {
// 		defer close(handleData)
// 		for i := range data {
// 			handleData <- data[i]
// 		}
// 	}
// 	handleData := make(chan int)
// 	go loopData(handleData)

// 	for num := range handleData {
// 		fmt.Println(num)
// 	}
// }

func main() {
	str := []byte("foobar")

	ch := make(chan byte, len(str))
	next := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < len(str); i++ {
		ch <- str[i]
	}

	close(ch)

	go func() {
		defer wg.Done()
		for {
			stop, ok := <-next
			if !ok {
				return
			}
			v, ok := <-ch
			if ok {
				fmt.Println("goroutine-1: ", string(v))
			} else {
				close(next)
				return
			}
			next <- stop
		}

	}()

	go func() {
		defer wg.Done()
		for {
			stop, ok := <-next
			if !ok {
				return
			}
			v, ok := <-ch
			if ok {
				fmt.Println("goroutine-2: ", string(v))
			} else {
				close(next)
				return
			}
			next <- stop
		}

	}()

	next <- struct{}{}
	wg.Wait()
}
