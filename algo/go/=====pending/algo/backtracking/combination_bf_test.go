package backtracking

import (
	"fmt"
	"testing"
)

func TestCombinationbf(t *testing.T) {

	fmt.Println("cobination bf n=3,k=3")
	a1 := combination_bf(3, 3)
	fmt.Println(a1)

	fmt.Println("combination recur n=3,k=3")
	a2 := combination_recur(5, 4)
	fmt.Println(a2)
}
