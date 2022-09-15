package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyCalendarTwo struct {
	m map[int]int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	arr := []string{}
	myC := Constructor()

	for scanner.Scan() {
		arr = strings.Split(scanner.Text(), " ")

		if len(arr) != 2 {
			fmt.Println("please enter 2 numbers")
			continue
		}
		start, _ := strconv.Atoi(arr[0])
		end, _ := strconv.Atoi(arr[1])

		fmt.Println(myC.Book(start, end))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{m: make(map[int]int)}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for i := start; i < end; i++ {
		if (this.m[i] + 1) > 3 {
			return false
		}
	}

	for i := start; i < end; i++ {
		this.m[i] += 1
	}

	fmt.Println(this.m)
	return true
}
