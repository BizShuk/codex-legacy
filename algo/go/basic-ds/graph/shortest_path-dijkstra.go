package graph

// [Pattern]: [Minimum Spanning Tree, MST] [Dijkstra] use min-heap/findMin from "Root" with bfs to reach all nodes from source
// Implementation: network delay time
// Use Greedy Algo to approach
// [Hint]: Node distance will be distance from source
// [Hint]: May/not work on negative edges, but definitely not work for negative edge/cycle

// [Ref]: Check [Prim] in graph/minimum_spanning_tree.go

// [Time Complexity]: O(ELogV), go through all edges from minimum one, LogV for poping out
// [Space Complexity]: O(V^2), store edges of each vertex (max number: V-1)

// [Data structure]: []distance (infinity for unvisited), map[node]edge, minheap(edges with neighbor of connected nodes)
func Dijkstra() {
	// 1. Create visited set and distance list

	// 2. Assign to every node a tentative distance value:
	// set it to zero for our initial node.
	// set to infinity for all other nodes.

	// 3. Pick up current node and consider all of its unvisited neighbors and calculate their tentative distances through the current node.
	// [Issue]: If consider visited nodes with better distance, it might cause infinity cycle with negative edges/cycle.

	// 4. Mark the current node as visited. A visited node will never be checked again (this is valid and optimal in connection with the behavior in step 6.: that the next nodes to visit will always be in the order of "smallest distance from initial node first"

	// 5. If the destination node was marked visited => There is a route between them.
	//    If the destination node's distance is Infinity => The node is not reachable
	//  then stop. The algorithm has finished.

	// 6. Otherwise, select the unvisited node that is marked with the smallest tentative distance, set it as the new current node, and go back to step 3.
}
