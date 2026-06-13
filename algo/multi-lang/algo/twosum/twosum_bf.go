package twosum

import (
	"sort"
)

// Brute Force
func twoSum_bf(nums []int, target int) []int {
	result := make([]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)

			}
		}
	}
	return result
}
