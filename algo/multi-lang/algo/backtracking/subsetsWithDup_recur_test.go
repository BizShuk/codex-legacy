package backtracking

import (
	"fmt"
	"testing"
)

func TestSubset2WithDup_recur(t *testing.T) {
	case1 := []int{0}
	case2 := []int{1, 2, 2}
	case3 := []int{1, 2, 2, 4}
	case4 := []int{1, 2, 2, 4, 4, 5}
	case5 := []int{1, 2, 2, 2, 4, 5, 5, 5}
	case6 := []int{4, 1, 0}

	exec(case1)
	exec(case2)
	exec(case3)
	exec(case4)
	exec(case5)
	exec(case6)

}

func exec(cases []int) {
	a := SubsetsWithDup_recur(cases)
	fmt.Println("case:", cases, "\nresult:", a)
}
