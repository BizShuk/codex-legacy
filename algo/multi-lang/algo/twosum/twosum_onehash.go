package twosum

func twoSum_hash(nums []int, target int) []int {
	list_nums := make(map[int]int)
	var result []int

	for i, v := range nums {
		if i2, ok := list_nums[target-v]; ok && i != list_nums[target-v] {
			result = append(result, i, i2)
			return result
		}
		list_nums[v] = i
	}
	return result
}
