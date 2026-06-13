package op

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	case1 := "abcdefgh"
	a1 := ReverseString(case1)
	fmt.Println(case1, a1)
	case2 := ""
	a2 := ReverseString(case2)
	fmt.Println(case2, a2)
	case3 := "a"
	a3 := ReverseString(case3)
	fmt.Println(case3, a3)
	/*
		fmt.Println("No Extra Space:")
		a1 = ReverseString_nospace(case1)
		fmt.Println(case1, a1)
		a2 = ReverseString_nospace(case2)
		fmt.Println(case2, a2)
		a3 = ReverseString_nospace(case3)
		fmt.Println(case3, a3)
	*/
}
