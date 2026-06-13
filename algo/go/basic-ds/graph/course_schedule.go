package graph

// [Pattern]: [Topology Sort] traverse graph, the vertex only visited when any vertex to this vertex visited

/*
   1 -> 2
   1 -> 3
   2 -> 3
   => 1 2 3  =x=> 1 3 2
*/

// [Variant]: [Topology Sort] Detect courses cycly - DFS
// [Hint]: cache(visited) for reducing process time
// [Hint]: trace stack for detecting if there is cycle in the same stack
// [Hint]: if visited and no cycle, any vertex connected to those vertexs will be no cycle too.

func CourseMapCycle(numCourses int, prerequisites [][]int) bool {
	courses := make([][]bool, 0)
	for i := 0; i < numCourses; i++ {
		courses = append(courses, make([]bool, numCourses))
	}

	for i := 0; i < len(prerequisites); i++ {
		preCourse, course := prerequisites[i][1], prerequisites[i][0]
		courses[course][preCourse] = true
	}

	var dfs func(courseI int, visited []bool, traceStack []bool) bool

	dfs = func(courseI int, visited []bool, traceStack []bool) bool {

		if traceStack[courseI] {
			return false
		}

		if visited[courseI] {
			return true
		}

		visited[courseI], traceStack[courseI] = true, true

		for i := 0; i < len(courses[courseI]); i++ {
			if !courses[courseI][i] {
				continue
			}

			if !dfs(i, visited, traceStack) {
				return false
			}

		}
		traceStack[courseI] = false

		return true
	}

	visited := make([]bool, numCourses)
	for i := 0; i < len(courses); i++ {
		if visited[i] {
			continue
		}
		traceStack := make([]bool, numCourses)
		if !dfs(i, visited, traceStack) {
			return false
		}
	}
	return true
}

// [Variant]: [Topology Sort] Detect courses cycly - BFS
// 1. Count all prerequisites for each vertex.
// 2. Put all 0 count vertext in queue for BFS
// 3. once reach, count--. If count == 0, push it into queue
// 4. if count == course count, no cycle
// 5. if count < course count, cycle detected
func CourseMapCycle_BFS(numCourses int, prerequisites [][]int) bool {
	pCount := make([]int, numCourses)
	edges := make([][]int, numCourses)

	f, t := 0, 0
	for i := 0; i < len(prerequisites); i++ {
		f, t = prerequisites[i][1], prerequisites[i][0]
		edges[f] = append(edges[f], t)
		pCount[t]++
	}

	queue, count := []int{}, 0
	for i := 0; i < numCourses; i++ {
		if pCount[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		f := queue[0]
		queue = queue[1:]

		t := 0
		for i := 0; i < len(edges[f]); i++ {
			t = edges[f][i]
			pCount[t]--
			if pCount[t] == 0 {
				queue = append(queue, t)
			}
		}

		count++
	}

	return count == numCourses
}
