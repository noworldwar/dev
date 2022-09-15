package main

import "fmt"

func main() {
	input := []int{9, 5, 1, 65, 7, 8, 2, 6, 3}
	result := BubbleSort(input)
	fmt.Println(result)
}

func BubbleSort(arr []int) []int {

	count := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			count++
		}

	}
	fmt.Printf("%d", count)
	return arr
}
