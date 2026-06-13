package hard

// [Question]: Find first missing positive number

// [Solution] 1: Sort it and go through it
// [Time Complexity]: O(NlogN)
// [Space Complexity]: O(1)

// [Solution] 2: Store data in hashmap and loop started from 1 until missing found
// [Time Complexity]: O(N)
// [Space Complexity]: O(N), only positive number in hashmap

// [Solution] 3: Eliminate
// [Time Complexity]: O(N)
// [Space Complexity]: O(N)

func FirstMissingPositive(nums []int) int {
	i, v, size := 0, 0, len(nums)

	for i < size {
		v = nums[i]
		if v <= 0 || v > size || v == nums[v-1] {
			i += 1
			continue
		}

		nums[i], nums[v-1] = nums[v-1], nums[i]
	}

	for i := 0; i < size; i++ {
		if i != nums[i]-1 {
			return i + 1
		}
	}
	return size + 1
}
