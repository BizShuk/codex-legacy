package op

import (
	"fmt"
	"testing"
)

func TestUglyNum(t *testing.T) {
	cases := 3
	a := UglyNum(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 21
	a = UglyNum(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 26
	a = UglyNum(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 30
	a = UglyNum(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 91361722
	a = UglyNum(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = -2147483648
	a = UglyNum(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 0
	a = UglyNum(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 1
	a = UglyNum(cases)
	fmt.Println("case:", cases, "result:", a)
}
