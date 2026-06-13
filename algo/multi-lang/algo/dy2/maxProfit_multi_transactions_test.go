package dy

import (
	"fmt"
	"testing"
)

func TestMaxProfit_multi_transactions(t *testing.T) {
	case1 := []int{1, 2, 3, 4, 1, 0, 2, -2}
	exec_multi_transactions(case1)
	case2 := []int{1, 2}
	exec_multi_transactions(case2)

	exec_multi_transactions(case8)
}

func exec_multi_transactions(cases []int) {
	a := MaxProfit_multi_transactions(cases)
	fmt.Println("Case:", cases, "\nresult:", a)

}
