package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, _ := strconv.ParseInt("一二三", 10, 64)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println(i)
}
