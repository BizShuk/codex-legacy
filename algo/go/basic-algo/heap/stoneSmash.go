package heap

import "container/heap"

// [Variant]: [Heap] Smash two biggest stones and find last remaining.
func LastStoneWeight(stones []int) int {

	maxHeap := &MaxHeap{}
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
	if x > y {
		return x - y
	}
	return y - x
}
