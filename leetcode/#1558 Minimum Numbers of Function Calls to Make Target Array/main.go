package main

import (
	"fmt"
)

func main() {
	// i := []int{0, 0, 0}
	j := []int{0}
	count := 0
	c1 := 0
	c2 := 0
	for {
		j, c2 = divideTwo(j)
		j, c1 = removeOne(j)
		count += c1 + c2
		if checkIfZero(j) == true {
			break
		}
	}

	fmt.Println(j)
	fmt.Println(count)
}

func checkIfZero(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum == 0
}

func removeOne(nums []int) ([]int, int) {
	count := 0
	for i, v := range nums {
		if v == 1 {
			fmt.Println("-1")
			nums[i]--
			count++
		}
	}
	return nums, count
}

func divideTwo(nums []int) ([]int, int) {
	new_s := []int{}
	count := 1
	for _, v := range nums {
		if v == 1 {
			return nums, 0
		}
		if v%2 == 0 {
			new_s = append(new_s, v/2)
		} else if v%2 == 1 {
			new_s = append(new_s, (v-1)/2)
			count += 1
			fmt.Println("+1")
		} else {
			return nums, 0
		}
	}

	if len(new_s) != len(nums) {
		return nums, 0
	}

	fmt.Println("/2")
	return new_s, count
}

func getModCount(nums []int) int {
	count := 0
	for _, v := range nums {
		count = count + (v % 2)
	}
	return count
}
