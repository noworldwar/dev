package main

import "fmt"

func main() {

	input := []int{3, 6, 9, 12}
	fmt.Println(longestArithSeqLength(input))

}

func longestArithSeqLength(nums []int) int {

	dp := make([]map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make(map[int]int)
	}

	now := 2
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			diff := nums[j] - v
			if prev, ok := dp[i][diff]; ok {
				dp[j][diff] = prev + 1
				if prev+1 > now {
					now = prev + 1
				}
			} else {
				dp[j][nums[j]-v] = 2
			}
		}
	}

	return now
}
