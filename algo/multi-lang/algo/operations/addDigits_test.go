package op

import (
	"fmt"
	"testing"
)

func TestAddDigits(t *testing.T) {
	case1 := 10
	a1 := AddDigits(case1)
	fmt.Println("case:", case1, "result:", a1)
	case2 := 1023123
	a2 := AddDigits(case2)
	fmt.Println("case:", case2, "result:", a2)
	case3 := 9
	a3 := AddDigits(case3)
	fmt.Println("case:", case3, "result:", a3)
	case4 := 512321312
	a4 := AddDigits(case4)
	fmt.Println("case:", case4, "result:", a4)
	case5 := 29
	a5 := AddDigits(29)
	fmt.Println("case:", case5, "result:", a5)
}

func TestAddDigits_math(t *testing.T) {
	case1 := 10
	a1 := AddDigits_math(case1)
	fmt.Println("case:", case1, "result:", a1)
	case2 := 1023123
	a2 := AddDigits_math(case2)
	fmt.Println("case:", case2, "result:", a2)
	case3 := 9
	a3 := AddDigits_math(case3)
	fmt.Println("case:", case3, "result:", a3)
	case4 := 512321312
	a4 := AddDigits_math(case4)
	fmt.Println("case:", case4, "result:", a4)
	case5 := 29
	a5 := AddDigits_math(29)
	fmt.Println("case:", case5, "result:", a5)
}
