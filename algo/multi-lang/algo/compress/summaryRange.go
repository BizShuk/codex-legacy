package summaryRange

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {

	var queue []int
	var result []string
	var last int

	for i, v := range nums {
		if i != 0 && last != v-1 {
			range_elem := resetQueue(queue)
			result = append(result, range_elem)
			queue = queue[:0]
		}

		queue = append(queue, v)
		last = v
	}
	range_elem := resetQueue(queue)
	result = append(result, range_elem)

	return result
}

func resetQueue(q []int) string {
	if len(q) == 0 {
		return ""
	}
	start, last := strconv.Itoa(q[0]), strconv.Itoa(q[len(q)-1])
	if start == last {
		return start
	}
	return fmt.Sprint(start, "->", last)
}
