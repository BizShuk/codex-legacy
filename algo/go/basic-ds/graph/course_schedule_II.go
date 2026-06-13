package graph

// [Variant]: [Topology Sort] Print course order
// [Hint]: Use counter for incoming edge. Once counter is zero, it's that vertex turn.
func CourseScheduleOrder(numCourses int, prerequisites [][]int) []int {

	pEdges := make([][]int, numCourses)
	pCount := make([]int, numCourses)
	courseSeq := make([]int, 0)

	f, t := 0, 0
	for i := 0; i < len(prerequisites); i++ {
		f, t = prerequisites[i][1], prerequisites[i][0]
		pEdges[f] = append(pEdges[f], t)
		pCount[t]++
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if pCount[i] == 0 {
			queue = append(queue, i)
			courseSeq = append(courseSeq, i)
		}
	}

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		for i := 0; i < len(pEdges[course]); i++ {
			nCourse := pEdges[course][i]
			pCount[nCourse]--
			if pCount[nCourse] == 0 {
				queue = append(queue, nCourse)
				courseSeq = append(courseSeq, nCourse)
			}
		}
	}

	if len(courseSeq) != numCourses {
		return []int{}
	}

	return courseSeq
}
