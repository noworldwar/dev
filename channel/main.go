package main

import (
	"fmt"
	"os"
	"time"
)

// func main() {
// 	count := 0
// 	for {
// 		go func() {
// 			fmt.Printf("hello %d \r\n", count)
// 			count++
// 		}()
// 		time.Sleep(time.Second * 2)
// 	}
// }

// func main() {
// count := 0

// for {
// 	go func() {
// 		fmt.Printf("hello %d \r\n", count)
// 		count++
// 	}()
// 	time.Sleep(time.Second * 2)
// }
// }

// func main() {

// 	var ch = make(chan string, 3)
// Loop:
// 	for {
// 		select {
// 		case ch <- "Got it":
// 			fmt.Println("Got value")
// 		case data := <-ch:
// 			fmt.Println("data: ", data)
// 		case <-time.After(time.Second * 3):
// 			fmt.Println("Request Time Out")
// 			break Loop
// 		}
// 	}

// }

// func main() {
// 	wg := sync.WaitGroup{}
// 	wg.Add(20)
// 	for i := 0; i < 20; i++ {
// 		go func(val int, wg *sync.WaitGroup) {
// 			fmt.Println("finisied job id: ", val)
// 			wg.Done()
// 		}(i, &wg)
// 	}

// 	wg.Wait()
// }

func main() {
	// ...create abort channel...
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("launch")
}
