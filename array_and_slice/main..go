package main

import (
	"fmt"
)

//array
// func main() {
// 	a := [2]string{"foo", "bar"}
// 	// var a [2]string

// 	a[0] = "foo"
// 	a[1] = "bar"

// 	fmt.Println(a[0]) //foo
// }

//slice
func main() {
	a := []string{"1", "2", "3", "4", "5", "6"}
	fmt.Println(a[0:1]) //[1]
	fmt.Println(a[1:1]) //[]
	fmt.Println(a[1:])  //[2,3,4,5,6]
	fmt.Println(a[:2])  //[1]
}
