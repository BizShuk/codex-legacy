// nums = [1,2,2] , assume it's sorted
// Solution:
// [
//
//	  [2],
//	  [1],
//		 [1,2,2],
//		 [2,2],
//		 [1,2],
//		 []
//
// ]
package backtracking

import (
	s "sort"

	"github.com/bizshuk/algo/pending/algo/sort"
)

func SubsetsWithDup_recur(nums []int) [][]int {
	s.Ints(nums)
	result_arr := [][]int{}

	// select n elements from 0 ~ length
	for size := 0; size <= len(nums); size++ {
		_sets := subset_recur([]int{}, nums, size)
		result_arr = append(result_arr, _sets...)
	}

	sort.Sort2Dcolumn(result_arr)
	return result_arr
}

func subset_recur(s []int, r []int, size int) [][]int {
	if len(s) == size {
		result_set := append([]int{}, s...)
		return [][]int{result_set}
	}

	result_arr := [][]int{}
	lastv := 0
	for i, v := range r {
		if i != 0 && lastv == v {
			continue
		}
		_s := append(s, v)
		_sets := subset_recur(_s, r[i+1:], size)
		result_arr = append(result_arr, _sets...)
		lastv = v
	}
	return result_arr
}
