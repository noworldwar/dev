package main

import "fmt"

type Tree struct {
	root *TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	inp := []int{-10, -3, 0, 5, 9}
	result := sortedArrayToBST(inp)
	fmt.Println(result)
}

func sortedArrayToBST(nums []int) *TreeNode {
	return nil
}

func checkOrder(nums []int) bool {
	return false
}
