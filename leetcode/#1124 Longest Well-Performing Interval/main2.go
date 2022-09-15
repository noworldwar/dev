package main

import "fmt"

func main() {
	s := []int{6, 6, 9, 9, 9, 0, 1, 3, 6, 7, 9, 9, 9}
	fmt.Println(longestWPI(s))
	// fmt.Println(longestWPI2(s))
}

func longestWPI(hours []int) (wpi int) {
	for i, v := range hours {
		if v > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}
	fmt.Println(hours)
	hash := map[int]int{}
	curr := 0
	for i, v := range hours {
		fmt.Println(curr, "+=", v)
		curr += v
		fmt.Println("put ", i, " to ", curr)
		hash[curr] = i
	}
	fmt.Println(hash)
	for k, v := range hash {
		if k > 0 && v+1 > wpi {
			wpi = v + 1
			fmt.Println("first ans: ", wpi)
		}
	}
	magic := 1
	for i, v := range hours {
		fmt.Println(magic, "+=", v)
		magic += v
		fmt.Println("wpi = ", hash[magic]-i)
		if hash[magic]-i > wpi {
			wpi = hash[magic] - i
		}
	}
	return wpi
}

func longestWPI2(nums []int) int {
	maxLen, count, hash := 0, 0, make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 8 {
			count++
		} else {
			count--
		}
		if count > 0 {
			maxLen = i + 1
		} else {
			fmt.Println(hash)
			if v, ok := hash[count-1]; ok {
				if i-v > maxLen {
					maxLen = i - v
				}
			}

			if _, ok := hash[count]; !ok {
				hash[count] = i
			}
		}
	}

	return maxLen
}
