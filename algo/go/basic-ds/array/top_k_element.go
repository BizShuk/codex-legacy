package array

// [Pattern]: [Bucket Sort] linear time complexity
// [Question]: Find top k element with linear time complexity
// [Hint]: Bucket sort
func TopKElement(nums []int, k int) []int {
	counter := make(map[int]int)

	for _, v := range nums {
		counter[v] += 1
	}

	bucket := make([][]int, len(nums)+1)
	for k, v := range counter {
		bucket[v] = append(bucket[v], k)
	}

	topKResult := []int{}
	for i := len(bucket) - 1; i >= 0; i-- { // [Time Complexity]: O(N), because only unique number in nums which is less than N or len(nums)
		for j := 0; j < len(bucket[i]); j++ { // this filter out the bucket don't have anything
			topKResult = append(topKResult, bucket[i][j])
			if len(topKResult) == k {
				break
			}
		}
	}
	return topKResult
}
