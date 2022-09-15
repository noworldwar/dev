package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sentence := ""
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	msg := make(chan string, 1)

	go func() {
		for {
			var s string
			fmt.Scan(&s)
			msg <- s
		}
	}()

loop:
	for {
		fmt.Println("please enter")
		select {
		case <-sigs:
			fmt.Println(sentence)
			break loop
		case s := <-msg:
			fmt.Println("Echoing: ", s)
			if s == "b" || s == "p" || s == "t" {
				sentence += s
			} else {
				fmt.Println("Please type 'b','p','t'")
			}

		}
	}
}
