func findPairs(nums []int, k int) int {
	cor := [][]int{}
	for n, _ := range nums {
		for i := n + 1; i < len(nums); i++ {
			if abs(nums[n]-nums[i]) == k {
				cor = append(cor, []int{nums[n], nums[i]})
			}

		}
	}
	return len(removeRepeat(cor))
}

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

func removeRepeat(nums [][]int) []string {
	keys := make(map[string]bool)
	val_str := ""
	rev_val_str := ""
	list := []string{}
	for _, val := range nums {
		val_str = strconv.Itoa(val[0]) + "," + strconv.Itoa(val[1])
		rev_val_str = strconv.Itoa(val[1]) + "," + strconv.Itoa(val[0])
		if val2 := keys[val_str]; !val2 {
			keys[val_str] = true
			keys[rev_val_str] = true
			list = append(list, val_str)
		}
	}

	return list
}