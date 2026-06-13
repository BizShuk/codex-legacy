package greedy

// [Question]: Find the town which you start with empty gas tank and can drive all towns back to start point.
// [Gas Station](https://leetcode.com/problems/gas-station/)

// [Time  Complexity]: O(n)
// [Space Complexity]: O(1)
func CanCompleteDriveCircuit(gas []int, cost []int) int {
	total, sum, sp, size := 0, 0, 0, len(gas)

	for i := 0; i < size; i++ {
		total += gas[i] - cost[i]
		sum += gas[i] - cost[i]

		if sum < 0 {
			sum, sp = 0, i+1
		}
	}

	if total < 0 {
		return -1
	}

	return sp
}
