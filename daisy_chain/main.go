package main

import "fmt"

func main() {
	const max = 100
	nums := xGenerator()
	number := <-nums

	for number <= max {
		fmt.Println(number)
		nums = filter(nums, number)
		number = <-nums
	}
}

func xGenerator() chan int {
	var ch chan int = make(chan int)

	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func filter(input chan int, num int) chan int {
	output := make(chan int)

	go func() {
		for {
			i := <-input

			if i%num != 0 {
				output <- i
			}
		}
	}()

	return output
}
