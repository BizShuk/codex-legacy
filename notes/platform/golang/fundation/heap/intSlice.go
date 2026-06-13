package heap

// TODO: figure out for promote field (IntSlice structu inside other struct)
// import (
// 	"container/heap"
// 	"fmt"
// 	"sort"
// )

// type IntHeap2 struct {
// 	sort.IntSlice
// }

// func (h *IntHeap2) Push(x interface{}) {
// 	// Push and Pop use pointer receivers because they modify the slice's length,
// 	// not just its contents.
// 	*h = append(*h, x.(int))
// }

// func (h *IntHeap2) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// func IntHeapMain2() {
// 	h := &IntHeap2{2, 1, 5}
// 	heap.Init(h)
// 	heap.Push(h, 3)
// 	fmt.Printf("minimum: %d\n", (*h)[0])
// 	for h.Len() > 0 {
// 		fmt.Printf("%d ", heap.Pop(h))
// 	}
// }
