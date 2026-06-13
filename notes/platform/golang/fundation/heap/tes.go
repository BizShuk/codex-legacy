package heap

import (
	"container/heap"
)

type KthLargest struct {
	minHeap *IntHeaptmp
	size    int
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeaptmp{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}

	return KthLargest{minHeap: h, size: k}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)
	for len(*this.minHeap) > this.size {
		heap.Pop(this.minHeap)
	}
	return this.minHeap.Peak()
}

type IntHeaptmp []int

func (h IntHeaptmp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeaptmp) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeaptmp) Len() int           { return len(h) }

func (h *IntHeaptmp) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeaptmp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *IntHeaptmp) Peak() int {
	return (*h)[0]
}
