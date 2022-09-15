package main

import (
	"math/rand"
	"time"
)

func main() {

}

type Solution struct {
	input []int
}

func Constructor(nums []int) Solution {
	return Solution{input: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.input
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	result := []int{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := r.Perm(len(this.input))
	for _, v := range perm {
		result = append(result, this.input[v])
	}
	return result
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
