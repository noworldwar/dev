package main

import "fmt"

func main() {
	arr := []int{1, 227, 257, 349, 431, 577, 599, 619, 709, 751, 787, 821, 1031, 1097, 1229, 1291, 1307, 1447, 1493, 1499, 1663, 1831, 2011, 2027, 2069, 2099, 2137, 2273, 2309, 2473, 2633, 2687, 2689, 2693, 2753, 2803, 2887, 2909, 2969, 3079, 3449, 3643, 3923, 3947, 4091, 4481, 4507, 4547, 4657, 4673, 4691, 4793, 5483, 5569, 5623, 5897, 5953, 6037, 6229, 6311, 6469, 6709, 6823, 6829, 6907, 7489, 7499, 7577, 7643, 7867, 7933, 7993, 8017, 8221, 8377, 8447, 8461, 9067, 9257, 9397, 9649, 9739, 9851, 10037, 10247, 10259, 10909, 11087, 11159, 11279, 11443, 11807, 12409, 12457, 12539, 12953, 12979, 13049, 13163, 13267, 13327, 13469, 13907, 13931, 14081, 14159, 14401, 14669, 14767, 15017, 15031, 15061, 15121, 15373, 15391, 15787, 15859, 15877, 15889, 15913, 15919, 16073, 16141, 16301, 16451, 17467, 17837, 17863, 17989, 18013, 18121, 18133, 18233, 18257, 18329, 18341, 18493, 18503, 18661, 18713, 18749, 18917, 19207, 19267, 19531, 19583, 19751, 19777, 19813, 19819, 19853, 20509, 20563, 20849, 21169, 21491, 21557, 21683, 21751, 22129, 22153, 22229, 22697, 22853, 22993, 23209, 23251, 23473, 23497, 23603, 23623, 23801, 23857, 23893, 24151, 24239, 24499, 24551, 24809, 24907, 25117, 25391, 25997, 26501, 26561, 26821, 26849, 26927, 27431, 27617, 27779, 28019, 28057, 28537, 28663, 29059, 29179, 29333, 29429, 29483}
	k := 1

	result := kthSmallestPrimeFraction(arr, k)
	fmt.Println(result[k-1])

	// fmt.Println(bigger([2]int{1, 2}, [2]int{3, 7}))

}

func kthSmallestPrimeFraction(arr []int, k int) [][]int {
	result := enumTheList(arr)
	// fmt.Println(result)
	return result
}

func enumTheList(arr []int) [][]int {
	result := [][]int{}
	for i := len(arr) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			result = append(result, []int{arr[j], arr[i]})
		}
	}
	return sortFromSmall(result)
}

func sortFromSmall(arr [][]int) [][]int {
	if len(arr) <= 1 {
		return arr
	}
	for i, _ := range arr {
		for j := i + 1; j < len(arr); j++ {
			if bigger(arr[i], arr[j]) {
				arr = swapArr(i, j, arr)
			}
		}
	}
	return arr
}

func bigger(a, b []int) bool {
	if a[1] == b[1] {
		return a[0] > b[0]
	} else {
		return a[0]*b[1] > b[0]*a[1]
	}
}

func swapArr(a, b int, arr [][]int) [][]int {
	arr[a][0], arr[a][1] = arr[a][1], arr[a][0]
	return arr
}
