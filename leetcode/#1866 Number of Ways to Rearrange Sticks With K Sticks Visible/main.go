package main

import "fmt"

func main() {
	fmt.Println(numWays(1000, 3))
}

func numWays(n, k int64) int64 {
	mod := int64(1000000007)
	dp := [1001][1001]int64{}
	if n == k {
		return 1
	}
	if k == 0 {
		return 0
	}
	if dp[n][k] == 0 {
		dp[n][k] = numWays(n-1, k-1) + numWays(n-1, k)*(n-1)%mod
	}
	return dp[n][k] % mod
}
