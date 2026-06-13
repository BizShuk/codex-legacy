package sort

import (
	"fmt"
	"sort"
)

var sortData []int = []int{9, 1, 8, 2, 7, 3, 6, 4, 5}

func SortInts() {
	sort.Ints(sortData) // [In-place sorting]
	fmt.Println("Sort.Ints", sortData)
	sort.Sort(sort.Reverse(sort.IntSlice(sortData))) // [In-place reverse sorting]
	fmt.Println("Sort.Reverse", sortData)
}
