package graph

import "sort"

// [Pattern]: [Minimum Spanning Tree, MST] [Kruskal]
// 1. Sort edges with weight
// 2. Choose the edge is minimum of remaining, can't point to self, if vertice connected, ignore as well.
// 3. Connect vertices
// [Hint]: [Diffrerence between Kruskal and Boruvkas] "shortest edge from each node", insted of "shortest edge from connected components"

// [Variant]: [Kruskal] [Disjoint Set Union Find] Find MST
// [Time Complexity]: O(ElogE)
func MinCostConnectPoints_Kruskal(points [][]int) int {
	psize := len(points)
	edges := make([][]int, 0) // [index of Node A, index of Node b, distance]

	for i := 0; i < psize; i++ { // O(V^2)
		for j := i + 1; j < psize; j++ {
			distance := CalDist(points[i][0], points[i][1], points[j][0], points[j][1])
			edges = append(edges, []int{i, j, distance})
		}
	}

	sort.Slice(edges, func(a, b int) bool { // O(ElongE)
		return edges[a][2] < edges[b][2]
	})

	parent := make([]int, psize)
	for i := 0; i < psize; i++ {
		parent[i] = i
	}
	var findParent func(x int) int
	findParent = func(x int) int {
		if parent[x] != x {
			parent[x] = findParent(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		px, py := findParent(x), findParent(y)
		if px == py {
			return false
		}
		parent[px] = py // can use size/rank to link
		return true
	}

	sum := 0

	for _, edge := range edges {
		if union(edge[0], edge[1]) {
			sum += edge[2]
		}
	}

	return sum

}
