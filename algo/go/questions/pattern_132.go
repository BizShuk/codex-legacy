package algo

// Tags: #monotonic_stack #stack

// Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i], nums[j] and nums[k]
// such that i < j < k and nums[i] < nums[k] < nums[j].
// Return true if there is a 132 pattern in nums, otherwise, return false.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: false
// Explanation: There is no 132 pattern in the sequence.

// Example 2:
// Input: nums = [3,1,4,2]
// Output: true
// Explanation: There is a 132 pattern in the sequence: [1, 4, 2].

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	min := make([]int, len(nums))
	min[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min[i-1] {
			min[i] = nums[i]
		} else {
			min[i] = min[i-1]
		}
	}

	stack := make([]int, 0)
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > min[i] {
			for len(stack) > 0 && stack[len(stack)-1] <= min[i] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 && stack[len(stack)-1] < nums[i] {
				return true
			}
			stack = append(stack, nums[i])
		}
	}

	return false
}
