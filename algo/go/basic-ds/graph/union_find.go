package graph

// [Pattern]: [Disjoint Set Union-Find] When adding new components, new one will merge with old one become a group. This can help to find specific component in which group.
// Basic concept is to merge groups and connect root of group B to root of group A

// [Ref]: https://cp-algorithms.com/data_structures/disjoint_set_union.html#search-for-connected-components-in-an-image

// [Solution] 1: Easy solution. Just connect each node by edge
// [Time Complexity]: O(n) for Find_set, total O(n) for all nodes
var parent []int = []int{}

func Make_set(v int) {
	parent[v] = v
}

func Find_set(v int) int {
	// [Solution] 1: Recursive
	if v == parent[v] {
		return v
	}

	return Find_set(parent[v])
	// [Solution] 1: Iterative
	// for v != parent[v] {
	// 	v = parent[v]
	// }
	// return v
}

func Union_sets(a, b int) {
	a = Find_set(a)
	b = Find_set(b)
	if a != b {
		parent[b] = a
	}
}

// [Variant]: [Disjoint Set Union-Find]
// [Hint]: [Path compression optimization]. When merging union, try to connect to the root of the other one. => Achieve "O(logN)" for Find_root.
// [Hint]: [Union by group size/rank]. When union, connect root of small-size group to root of large-size group. Rank is simuliar to group size. [Time Complexity] will reach O(α(N)) α is the inverse Ackermann function
func FindRedundantConnection(edges [][]int) []int {
	parent := make([]int, 1001)
	size := make([]int, 1001)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	for _, edge := range edges {
		p1, p2 := FindRoot(parent, edge[0]), FindRoot(parent, edge[1])

		if p1 == p2 {
			return []int{edge[0], edge[1]}
		}

		// [Solution]: Group Size approach
		if size[p2] > size[p1] {
			p1, p2 = p2, p1
		}

		// [Solution]: Index approach
		// if p2 > p1 {
		// 	p1, p2 = p2, p1
		// }

		parent[p2] = p1
		size[p2] += size[p1]
	}

	return []int{}
}

func FindRoot(parent []int, node int) int {
	for parent[node] != node {
		node = parent[node]
	}
	return node
}

// [Variant]: [Disjoint Set Union-Find] Use index as group
// [Hint]: [Linking by index]. Simuliar approach to Union by group size, but give a index number to group, and use the index to compare. Theoratically, the speed should be around with Union by group size. In practical, it's slower than by group size.
type UnionDS []int // value can be group index/root index/...

func (u UnionDS) Find(x, y int) bool {
	return u[x] == u[y]
}

func (u UnionDS) Union(x, y int) bool {
	for i := 0; i < len(u); i++ {
		if u[i] == u[y] {
			u[i] = u[x]
		}
	}
	return true
}
