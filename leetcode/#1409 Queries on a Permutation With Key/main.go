func processQueries(queries []int, m int) []int {
	p := []int{}
	result := []int{}
	for i := 0; i < m; i++ {
		p = append(p, i+1)
	}

	for n := range queries {
		pos := findPlace(p, queries[n])
		if pos != -1 {
			result = append(result, pos)
			p = putFront(p, p[pos])
		}
	}

	return result
}

func findPlace(queries []int, num int) int {
	for i := 0; i < len(queries); i++ {
		if queries[i] == num {
			return i
		}
	}
	return -1
}

func putFront(queries []int, num int) []int {
	output := []int{}
	if queries[0] == num {
		return queries
	}
	if queries[len(queries)-1] == num {
		output = append([]int{num}, queries[:len(queries)-1]...)
		return output
	}

	for n := range queries {
		if queries[n] == num {
			output = append([]int{num}, append(queries[:n], queries[n+1:]...)...)
			return output
		}
	}
	return nil
}