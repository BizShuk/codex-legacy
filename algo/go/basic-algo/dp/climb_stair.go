package dp

import "github.com/bizshuk/algo/util"

// How many ways to climb stair? maximum two steps each time
func ClimbStairs(steps int) int {
	if steps <= 2 {
		return steps
	}

	s1, s2 := 1, 2

	for i := 3; i <= steps; i++ {
		s1, s2 = s2, s1+s2
	}
	return s2
}

// [Variant]: [Simple DP] Find minimum priority/cost to climb stairs
func ClimbStairsWithCost(costs []int) int {
	c1, c2 := 0, 0

	for _, cost := range costs {
		tmp := util.Min(c1, c2) + cost
		c1, c2 = c2, tmp
	}

	return util.Min(c1, c2)
}
