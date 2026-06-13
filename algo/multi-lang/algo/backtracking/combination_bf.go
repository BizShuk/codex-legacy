package backtracking

// select different k element in {1..n}
func combination_bf(n, k int) [][]int {

	if k > n {
		return [][]int{}
	}
	var result [][]int
	start := 1
	tmp := make([]int, k)
	for i := 0; i < k; i++ {
		tmp[i] = i + 1
	}
	for {
		if n < start+k-1 {
			break
		}
		i := 0
		for i < k {
			if tmp[i] > n {
				if i == 0 {
					break
				}
				i = i - 1
				tmp[i] = tmp[i] + 1
				for j := i + 1; j < k; j++ {
					tmp[j] = tmp[j-1] + 1
				}
				continue
			}
			i = i + 1
		}

		if tmp[0] > n-k+1 {
			break
		}
		tmp1 := make([]int, k)
		copy(tmp1, tmp)
		result = append(result, tmp1)
		tmp[k-1] = tmp[k-1] + 1
		start = tmp[0]
	}
	return result
}
