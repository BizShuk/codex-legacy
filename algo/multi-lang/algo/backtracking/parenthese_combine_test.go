package backtracking

import (
	"fmt"
	"testing"
)

func TestParenthese_Combine(t *testing.T) {
	case1 := 1
	case2 := 2
	case3 := 3
	case4 := 4
	case5 := 5

	a1 := Parenthese_Combine(case1)
	fmt.Println("Len:", len(a1), "\nResult:", a1)

	a2 := Parenthese_Combine(case2)
	fmt.Println("Len:", len(a2), "\nResult:", a2)

	a3 := Parenthese_Combine(case3)
	fmt.Println("Len:", len(a3), "\nResult:", a3)

	a4 := Parenthese_Combine(case4)
	fmt.Println("Len:", len(a4), "\nResult:", a4)

	a5 := Parenthese_Combine(case5)
	fmt.Println("Len:", len(a5), "\nResult:", a5)
}
