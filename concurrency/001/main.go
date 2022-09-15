package main

import (
	"fmt"
)

func main() {
	result := consumer()
	for v := range result {
		fmt.Println(v)
	}
	fmt.Println()
}

func produce(nums ...int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, v := range nums {
			out <- v
		}
	}()

	return out
}

//         (inCh <- chan int) 也可以
func square(inCh chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range inCh {
			out <- v * v
		}
	}()

	return out
}

func consumer() chan int {

	in := produce(1, 2, 3, 5, 4, 9)
	ch := square(in)

	return ch
}
