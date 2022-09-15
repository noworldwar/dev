package main

import (
	"fmt"
	"time"
)

func main() {

	//for loop
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	// j := 0
	// for ; j < 5; j++ {
	// 	fmt.Println(j)
	// }

	//foreach
	// data := [5]string{"foo", "bar", "foobar", "foobar2000"}
	// for index, value := range data {
	// 	fmt.Println(index, value)
	// }

	// for index := range data {
	// 	fmt.Printf("%d", index)
	// }

	// for _, i := range data {
	// 	fmt.Println(i)
	// }
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "One"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "Two"
	}()

Loop:
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		default:
			fmt.Println("Waiting...")
			time.Sleep(time.Second * 3)
			continue Loop
		}
	}
}
