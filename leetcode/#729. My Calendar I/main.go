package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyCalendar struct {
	Activity [][]int64
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
		start, _ := strconv.ParseInt(arr[0], 10, 32)
		end, _ := strconv.ParseInt(arr[1], 10, 32)

		fmt.Println(myC.Book(start, end))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int64, end int64) bool {
	for _, v := range this.Activity {
		if v[0] < end && start < v[1] {
			return false
		}
	}
	this.Activity = append(this.Activity, []int64{start, end})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
