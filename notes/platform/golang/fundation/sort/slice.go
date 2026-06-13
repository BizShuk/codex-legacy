package sort

import (
	"fmt"
	"sort"
)

// [Pattern]: [Go][Go Anonymous struct] Sort multiple fields
func SortSlice() {
	fmt.Println("Hello, 世界")

	x := []struct {
		Name string
		Val  int
	}{
		{"Poor", 3},
		{"Fair", 2},
		{"Good", 0},
		{"Excellent", 3},
		{"Elite", 0},
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i].Val < x[j].Val
	})

	fmt.Println(x)
}
