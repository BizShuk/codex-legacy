package array

import "github.com/bizshuk/algo/util"

// [Question]: Given array with 1 ~ n value in n+1 size slice. Find duplicate.

func FindDuplicate(nums []int) int {
	numMap := map[int]bool{}

	for _, num := range nums {
		if numMap[num] {
			return num
		}
		numMap[num] = true
	}
	return -1
}

// [Question]: Given array with 1 ~ n value in n+1 size slice, find duplicate,
// [Time Complexity]: O(N)
// [Space Complexity]: O(1)
// [Hint]: Leverage index to find the slot for each [1,n] value
// N+1 elements with [1,n] value => one element must be duplicate
// step by step:
//     [ 4, 3, 2, 1, 4]
//     [ 4, 3, 2,-1, 4]
//     [ 4, 3,-2,-1, 4]
//     [ 4,-3,-2,-1, 4]
//     [-4,-3,-2,-1, 4]
//     return 4

func FindDuplicate_b(arr []int) int {

	for i := 0; i < len(arr); i++ {
		pnum := util.Abs(arr[i])
		if arr[pnum-1] > 0 {
			arr[pnum-1] = -arr[pnum-1]
		} else {
			return util.Abs(arr[i])
		}
	}
	return -1
}
