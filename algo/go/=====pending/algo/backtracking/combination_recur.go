package backtracking

// recursive version
func combination_recur(n, k int) [][]int {
	var result [][]int
	for i := 1; i <= n-k+1; i++ {
		selected := cb_recur([]int{i}, i+1, n, k-1)
		result = append(result, selected...)
	}
	return result

}

func cb_recur(t []int, start, end, size int) [][]int {
	//fmt.Println(t, start, end, size)
	if size == 0 {
		selected := make([]int, len(t))
		copy(selected, t)
		return [][]int{selected}
	}
	var result [][]int
	for i := start; i <= end; i++ {
		tmp_combination := append(t, i)
		selected := cb_recur(tmp_combination, i+1, end, size-1)
		result = append(result, selected...)
	}
	return result
}
