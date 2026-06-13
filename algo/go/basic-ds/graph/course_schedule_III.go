package graph

import (
	"container/heap"
	"sort"

	h "github.com/bizshuk/algo/heap"
	"github.com/bizshuk/algo/util"
)

// [Variant]: [Topology Sort][Dynamic Programming] find max number of courses can be finished
// [Solution]: knapsack solution (DP) aka climb stair,
func ScheduleCourse_2(courses [][]int) int {

	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] <= courses[j][1]
	})

	dp := make([]int, courses[len(courses)-1][1]+1)

	for _, course := range courses {
		// [Hint]:
		for lastDay := course[1]; lastDay-course[0] >= 0; lastDay-- {
			dp[lastDay] = util.Max(dp[lastDay], dp[lastDay-course[0]]+1)
		}
	}

	result := 0

	for i := 0; i <= courses[len(courses)-1][1]; i++ {
		result = util.Max(result, dp[i])
	}

	return result

}

// [Variant]: [Topology Sort]
// [Solution]: use Priority Queue to pop max duration course taken
func ScheduleCourse_3(courses [][]int) int {

	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] <= courses[j][1]
	})

	maxHeap := &h.MaxHeap{}
	time := 0

	for _, course := range courses {
		if time+course[0] <= course[1] {
			heap.Push(maxHeap, course[0])
			time += course[0]
			continue
		}

		if maxHeap.Len() > 0 && maxHeap.Top() > course[0] {
			time += -maxHeap.Top() + course[0]
			heap.Pop(maxHeap)
			heap.Push(maxHeap, course[0])
		}
	}

	return maxHeap.Len()
}

// [Variant]: [Topology Sort] find max number of courses can be finished
// [Solution]: use start time instead of deadline, and do permuataion algo
// [Result]: Time Exceeds limit
func ScheduleCourse_1(courses [][]int) int {
	latestStartDate := make([][]int, 0)
	numCourses := len(courses)
	for i := 0; i < numCourses; i++ {
		d, l := courses[i][0], courses[i][1]
		if d > l {
			continue
		}
		latestStartDate = append(latestStartDate, []int{l - d, d})
	}
	numCourses = len(latestStartDate)
	// do permutation with heap
	day, count, maxCount := 0, 0, 0

	var SwapSeq func(idx int)

	SwapSeq = func(idx int) {
		idx++

		if idx >= numCourses || maxCount == len(latestStartDate) {
			return
		}

		for i := idx; i < numCourses; i++ {
			if day > latestStartDate[i][0] {
				continue
			}

			latestStartDate[i], latestStartDate[idx] = latestStartDate[idx], latestStartDate[i]
			day += latestStartDate[idx][1]
			count += 1
			if count > maxCount {
				maxCount = count
			}

			SwapSeq(idx)
			count -= 1
			day -= latestStartDate[idx][1]

			latestStartDate[i], latestStartDate[idx] = latestStartDate[idx], latestStartDate[i]
		}

	}

	SwapSeq(-1)
	return maxCount
}
