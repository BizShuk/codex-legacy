package graph

import (
	"math"
)

// [Variant]: [Bellman-Ford] Network Delay
func NetworkDelayTime_Bellman_Ford(times [][]int, n int, k int) int {
	vSize, maxLength := n, math.MaxInt

	distance := make([]int, vSize)
	for i := 0; i < vSize; i++ {
		distance[i] = maxLength
	}
	distance[k-1] = 0

	for i := 0; i < n; i++ {

		for _, e := range times {
			u, v, w := e[0], e[1], e[2]
			if distance[u-1] == maxLength {
				continue
			}

			if distance[v-1] > distance[u-1]+w {
				distance[v-1] = distance[u-1] + w
			}

		}
	}

	time := 0
	for _, d := range distance {
		if d == maxLength {
			return -1
		}
		if time < d {
			time = d
		}
	}

	return time
}

// [Variant]: [Dijkstra] Network Delay with clean Python
// class Solution:
//     def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
//         seen=set()
//         edges = defaultdict(list)

//         for s,t, w in times:
//             edges[s].append((t,w))

//         heap=[(0,k)]
//         heapq.heapify(heap)
//         restime=0

//         while heap:
//             time, node = heapq.heappop(heap)
//             if node in seen:
//                 continue
//             restime=time
//             seen.add(node)

//             for dn, dt in edges[node]:
//                 if dn not in seen:
//                     heapq.heappush(heap,(time+dt, dn))
//         if len(seen)==n:
//             return restime
//         return -1

// [Variant]: [Dijkstra] Network Delay, not optimized with Heap (worst case scenario)
// [Result]: Time Exceeding
func NetworkDelayTime_Dijkstra(times [][]int, n int, k int) int {
	time := 0
	vSize, eSize, maxLength := n, len(times), n*101

	edges := make([][]int, vSize)
	for i := range edges {
		for j := 0; j < vSize; j++ {
			edges[i] = append(edges[i], -1)
		}

	}

	for i := 0; i < eSize; i++ {
		u, v, w := times[i][0], times[i][1], times[i][2]
		edges[u-1][v-1] = w
	}

	queue := []int{k - 1}

	distance := make([]int, vSize)
	for i := 0; i < vSize; i++ {
		distance[i] = maxLength
	}
	distance[k-1] = 0

	for len(queue) > 0 { // dijkstra = calculate all distances from root
		v := queue[0]
		queue = queue[1:]

		for i := 0; i < len(edges[v]); i++ {

			w := edges[v][i]
			if w == -1 || i == k-1 {
				continue
			}

			if distance[i] < distance[v]+w {
				continue
			}

			distance[i] = distance[v] + w
			queue = append(queue, i)
		}
	}

	for _, v := range distance {
		if v == maxLength {
			return -1
		}
		if time < v {
			time = v
		}
	}

	return time
}
