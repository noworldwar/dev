package main

import (
	"fmt"
)

func main() {

	// a := []int{}
	a := make([]int, 3, 5)

	fmt.Println(a)

	a = append(a, 1, 2, 3, 4, 4, 3, 6, 6, 5, 6, 7)

	fmt.Println(a)
	fmt.Println(removeDuplicateValues(a))
}

//Remove repeated element from array
func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
