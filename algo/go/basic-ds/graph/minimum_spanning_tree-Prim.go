package graph

// [Pattern]: [Minimum Spanning Tree, MST] [Prim] MST for weighted undirected graph
//
//	Pick one node and choose shortest edge from it, simular to Dijkstra's Algo "from the root", but find closer node "from the tree"
//
// [Time Complexity]: O(V^2)
// [Space Complexity]: O(V^2)

func MinCostConnectPoints(points [][]int) int {
	psize := len(points)

	distance := make([]int, psize) // [Note]: -1 已跑過, 目前已知到達i之最小距離
	for i := range distance {
		distance[i] = 0
	}

	cutpoint, sum := 0, -1
	for cutpoint >= 0 {
		distance[cutpoint] = -1
		minedge, nextpoint := -1, -1
		x, y := points[cutpoint][0], points[cutpoint][1]

		for i := range points {
			if distance[i] < 0 { // have connected
				continue
			}

			d := CalDist(x, y, points[i][0], points[i][1])

			if distance[i] == 0 || distance[i] > d { // Update distance if not update before or new one is shorter
				distance[i] = d
			}

			if minedge == -1 || minedge > distance[i] { // update minedge to find next point, minedge is to compare with distances from current node and explored distance from previous
				minedge = distance[i]
				nextpoint = i
			}

		}
		sum += minedge
		cutpoint = nextpoint

	}

	return sum
}
