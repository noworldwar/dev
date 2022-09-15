package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyCalendarTwo struct {
	once  [][2]int
	twice [][2]int
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
	return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {

	MAX := func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		} else {
			return i2
		}
	}
	MIN := func(i1, i2 int) int {
		if i1 < i2 {
			return i1
		} else {
			return i2
		}
	}
	// check twice pool
	for _, twi := range this.twice {
		if twi[0] < end && twi[1] > start {
			return false
		}
	}

	// check once pool
	for _, onc := range this.once {

		s := MAX(onc[0], start)
		e := MIN(onc[1], end)
		if s < e {
			this.twice = append(this.twice, [2]int{s, e})
		}

	}

	this.once = append(this.once, [2]int{start, end})

	return true
}
