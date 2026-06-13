package graph

// [Variant]: [Clone LinkedList]  Clone graph

// [Solution]: [DFS] Clone graph
// [Tip]: Use dfs to return new node of current node
func CloneGraph_dfs(node *Node) *Node {
	if node == nil {
		return nil
	}
	newNodeMap := map[*Node]*Node{}

	var dfs func(node *Node) *Node
	// func( oldNode ) newNode
	dfs = func(node *Node) *Node {
		if newNodeMap[node] != nil {
			return newNodeMap[node]
		}

		copyNode := &Node{Val: node.Val, Neighbors: []*Node{}}
		newNodeMap[node] = copyNode
		for _, next := range node.Neighbors {
			copyNode.Neighbors = append(copyNode.Neighbors, dfs(next))
		}
		return copyNode
	}

	return dfs(node)
}

// [Solution]: [BFS] Clone graph
func CloneGraph_bfs(node *Node) *Node {
	if node == nil {
		return nil
	}
	newNodeMap := map[*Node]*Node{}

	nodeQ := []*Node{node}
	// O(V)
	for len(nodeQ) > 0 {
		newNodeQ := []*Node{}
		for _, node := range nodeQ {
			newNodeMap[node] = &Node{Val: node.Val, Neighbors: []*Node{}}
			for _, next := range node.Neighbors {

				if newNodeMap[next] != nil {
					continue // walked
				}
				newNodeQ = append(newNodeQ, next)
			}
		}
		nodeQ = newNodeQ
	}
	// O(VE), Walk through all nodes (V)  and Visit all neighbors (E)
	for oldNode, newNode := range newNodeMap {
		for _, oldNext := range oldNode.Neighbors {
			newNext := newNodeMap[oldNext]
			newNode.Neighbors = append(newNode.Neighbors, newNext)
		}
	}

	return newNodeMap[node]
}
