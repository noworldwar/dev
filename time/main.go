package main

import (
	"fmt"
	"time"
)

func main() {
	//Get current time
	now := time.Now()
	fmt.Println("(Time)Current time: ", now)

	//Number of seconds since January 1, 1970 UTC
	unix_time := now.Unix()
	fmt.Println("(int64)Current time: ", unix_time)

	//Number of nanoseconds since January 1, 1970 UTC
	fmt.Println("(int64)Current time: ", now.UnixNano())

	//Convert timestamp to string
	int64time := time.Unix(unix_time, 0)
	strtime := int64time.Format("2006-01-02 03:04:05 PM")
	fmt.Println("(String)Current time: ", strtime)

	//Convert string to timestamp
	timestp, _ := time.Parse("2006-01-02 03:04:05 PM", strtime)
	fmt.Println("(Time)Current time: ", timestp)
}
