package hide

import "container/heap"

func lastStoneWeight(stones []int) int {

	maxHeap := &IntHeap{}
	heap.Init(maxHeap)

	for _, weight := range stones {
		heap.Push(maxHeap, weight)
	}

	for len(*maxHeap) > 1 {
		x, y := heap.Pop(maxHeap), heap.Pop(maxHeap)
		result := smash(x.(int), y.(int))
		heap.Push(maxHeap, result)
	}

	return (*maxHeap)[0]

}

func smash(x, y int) int {
	if x == y {
		return 0
	}

	if x > y {
		return x - y
	}
	return y - x
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
