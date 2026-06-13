package graph

// [Variant]: [Topology Sort][Floy-warshall] check whether one course is prerequisite of the other
// [Algo]: [Floy-warshall] This algorithm works for both the directed and undirected weighted graphs. But, it does not work for the graphs with negative cycles (where the sum of the edges in a cycle is negative).
func CoursePrerequisiteCheck(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	reachableMap := make([][]bool, 0)
	for i := 0; i < numCourses; i++ {
		reachableMap = append(reachableMap, make([]bool, numCourses))
	}

	for i := 0; i < len(prerequisites); i++ {
		u, v := prerequisites[i][0], prerequisites[i][1]
		reachableMap[u][v] = true
	}

	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				reachableMap[i][j] = reachableMap[i][j] || reachableMap[i][k] && reachableMap[k][j]
			}
		}
	}

	result := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		u, v := queries[i][0], queries[i][1]
		result[i] = reachableMap[u][v]
	}

	return result
}
