package main

import (
	"container/heap"
	"fmt"
)

type Fraction struct {
	Val       float64
	NumIndex  int // numberator index
	DenoIndex int // denominator index
}

type minHeap []Fraction

func main() {
	// arr := []int{1, 227, 257, 349, 431, 577, 599, 619, 709, 751, 787, 821, 1031, 1097, 1229, 1291, 1307, 1447, 1493, 1499, 1663, 1831, 2011, 2027, 2069, 2099, 2137, 2273, 2309, 2473, 2633, 2687, 2689, 2693, 2753, 2803, 2887, 2909, 2969, 3079, 3449, 3643, 3923, 3947, 4091, 4481, 4507, 4547, 4657, 4673, 4691, 4793, 5483, 5569, 5623, 5897, 5953, 6037, 6229, 6311, 6469, 6709, 6823, 6829, 6907, 7489, 7499, 7577, 7643, 7867, 7933, 7993, 8017, 8221, 8377, 8447, 8461, 9067, 9257, 9397, 9649, 9739, 9851, 10037, 10247, 10259, 10909, 11087, 11159, 11279, 11443, 11807, 12409, 12457, 12539, 12953, 12979, 13049, 13163, 13267, 13327, 13469, 13907, 13931, 14081, 14159, 14401, 14669, 14767, 15017, 15031, 15061, 15121, 15373, 15391, 15787, 15859, 15877, 15889, 15913, 15919, 16073, 16141, 16301, 16451, 17467, 17837, 17863, 17989, 18013, 18121, 18133, 18233, 18257, 18329, 18341, 18493, 18503, 18661, 18713, 18749, 18917, 19207, 19267, 19531, 19583, 19751, 19777, 19813, 19819, 19853, 20509, 20563, 20849, 21169, 21491, 21557, 21683, 21751, 22129, 22153, 22229, 22697, 22853, 22993, 23209, 23251, 23473, 23497, 23603, 23623, 23801, 23857, 23893, 24151, 24239, 24499, 24551, 24809, 24907, 25117, 25391, 25997, 26501, 26561, 26821, 26849, 26927, 27431, 27617, 27779, 28019, 28057, 28537, 28663, 29059, 29179, 29333, 29429, 29483}
	arr := []int{1, 2, 3, 5, 7}
	k := 3

	result := kthSmallestPrimeFraction(arr, k)
	fmt.Println(result)

	// fmt.Println(bigger([2]int{1, 2}, [2]int{3, 7}))

}

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m minHeap) Less(i, j int) bool {
	return m[i].Val < m[j].Val
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(Fraction))
}

func (m *minHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	old = old[:n-1]
	*m = old
	return x
}

/*
1/5 < 1/3 < 1/2
2/5 < 2/3
3/5
*/
func kthSmallestPrimeFraction(A []int, K int) []int {
	m := &minHeap{}
	heap.Init(m)

	// push first fractions into the heap.
	// we will push [ 1/5, 2/5, 3/5 ]
	n := len(A)
	for i := 0; i < n-1; i++ {
		f := Fraction{float64(A[i]) / float64(A[n-1]), i, n - 1}
		heap.Push(m, f)
	}
	result := []int{}
	// pop k elements
	for {
		f := heap.Pop(m)
		temp := f.(Fraction)
		fmt.Println("temp.NumIndex: ", temp.NumIndex)
		fmt.Println("temp.DenoIndex: ", temp.DenoIndex)
		K--
		if K == 0 {
			result = append(result, A[temp.NumIndex])
			result = append(result, A[temp.DenoIndex])
			return result
		}
		// push next fraction from poped chain
		nextNum := temp.NumIndex
		nextDeno := temp.DenoIndex - 1
		if nextNum < nextDeno {
			nextF := Fraction{float64(A[nextNum]) / float64(A[nextDeno]), nextNum, nextDeno}
			heap.Push(m, nextF)
		}
	}
	// return nil
}
