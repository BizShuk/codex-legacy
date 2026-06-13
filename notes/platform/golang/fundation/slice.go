package fundation

import "fmt"

// [Tip]: [Go] Unawared copy array
// Pass an array to function will result a new copy of array

// [Tip]: [Go] Slice memory leak
// slice an array/slice will cause underlying array kept unnecessary as long as slice keep.

func BasicSlice() {
	// s1 := []int{} // way of create slice
	s2 := make([]int, 0, 10)
	fmt.Printf("len(s2): %v, cap(s2): %v \n", len(s2), cap(s2))
	fmt.Println("append 1, 2,3 to s2")
	s2 = append(s2, 1, 2, 3)
	fmt.Printf("len(s2): %v, cap(s2): %v \n", len(s2), cap(s2))
}

// [Notice]: [Go] Slice leak element // TODO: find the test case from previous
