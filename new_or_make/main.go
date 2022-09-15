package main

import (
	"fmt"
	"sync"
	"time"
)

type user struct {
	name string
	age  int
	lock sync.Mutex
}

func main() {
	// u := new(user)

	// u.lock.Lock()
	// u.name = "hihi"
	// u.lock.Unlock()

	// fmt.Println(u)

	// var i *int
	// i = new(int)
	// *i = 10
	// fmt.Println(*i)

	ch := make(chan int)

	ch2 := make(chan string)

	ch3 := make(chan string)

	go func(ch chan int) {
		for i := 0; i < 26; i++ {
			ch <- i
		}
	}(ch)

	go func(ch2 chan string) {
		for i := 65; i < 91; i++ {
			ch2 <- string(i)
		}
	}(ch2)

	go func(ch3 chan string) {
		time.Sleep(time.Second * 10)
		ch3 <- "stop"
	}(ch3)

	for {
		select {
		case msg1 := <-ch:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case msg3 := <-ch3:
			fmt.Println(msg3)
			return
		default:
			fmt.Println("Waiting...")
			time.Sleep(time.Second * 1)
		}

	}

	close(ch)
	close(ch2)
	close(ch3)
}
