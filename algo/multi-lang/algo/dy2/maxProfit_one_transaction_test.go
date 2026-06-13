package dy

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	case1 := []int{1, 2, 3, 4, 1, 0, 2, -2}
	exec(case1)
	case2 := []int{1, 2}
	exec(case2)
}

func exec(cases []int) {
	a := MaxProfit1(cases)
	fmt.Println("Case:", cases, "\nresult:", a)

}
