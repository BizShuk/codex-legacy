package graph

import "math"

// Edge represents a directed edge with weight (may be negative).
type Edge struct {
	From, To string
	Weight   int
}

// BFResult holds the shortest-path computation result.
type BFResult struct {
	Distances   map[string]int
	Predecessor map[string]string
	HasNegCycle bool
	NegCycle    []string
}

// BellmanFord finds shortest path from `start` to all reachable nodes.
//
// [Pattern]: shortest path with negative-weight edges
// [Time Complexity]: O(V*E)
// [Space Complexity]: O(V + E)
//
// Note: Result only contains nodes reachable from `start`.
func BellmanFord(start string, nodes []string, edges []Edge) *BFResult {
	// Step 1: init distances to +Inf, predecessor to nil
	dist := make(map[string]int, len(nodes))
	pred := make(map[string]string, len(nodes))
	for _, n := range nodes {
		dist[n] = math.MaxInt32
	}
	dist[start] = 0

	// Step 2: relax all edges |V|-1 times
	for i := 0; i < len(nodes)-1; i++ {
		updated := false
		for _, e := range edges {
			if dist[e.From] == math.MaxInt32 {
				continue
			}
			newDist := dist[e.From] + e.Weight
			if newDist < dist[e.To] {
				dist[e.To] = newDist
				pred[e.To] = e.From
				updated = true
			}
		}
		// Early termination: a no-op pass means we've converged
		if !updated {
			break
		}
	}

	// Step 3: negative-cycle detection — one more relaxation pass
	for _, e := range edges {
		if dist[e.From] == math.MaxInt32 {
			continue
		}
		if dist[e.From]+e.Weight < dist[e.To] {
			// Cycle reachable. Walk back |V| steps to land inside the cycle,
			// then traverse predecessors until we revisit the start vertex.
			u := e.To
			for i := 0; i < len(nodes); i++ {
				u = pred[u]
			}
			cycle := []string{u}
			v := pred[u]
			for v != u {
				cycle = append(cycle, v)
				v = pred[v]
			}
			return &BFResult{
				Distances:   dist,
				Predecessor: pred,
				HasNegCycle: true,
				NegCycle:    cycle,
			}
		}
	}

	return &BFResult{
		Distances:   dist,
		Predecessor: pred,
		HasNegCycle: false,
	}
}
