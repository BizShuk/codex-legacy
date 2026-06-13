package op

import (
	"fmt"
	"testing"
)

func TestPowOfTwo(t *testing.T) {
	cases := 0
	a := PowOfTwo(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 1
	a = PowOfTwo(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 2
	a = PowOfTwo(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 10
	a = PowOfTwo(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 200
	a = PowOfTwo(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 1024
	a = PowOfTwo(cases)
	fmt.Println("case:", cases, "result:", a)
}

func TestPowOfThree(t *testing.T) {
	cases := 0
	a := PowOfThree(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 1
	a = PowOfThree(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 3
	a = PowOfThree(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 10
	a = PowOfThree(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 243
	a = PowOfThree(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 1024
	a = PowOfThree(cases)
	fmt.Println("case:", cases, "result:", a)
}

func TestPowOfFour_math(t *testing.T) {
	cases := 0
	a := PowOfFour_math(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 1
	a = PowOfFour_math(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 4
	a = PowOfFour_math(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 10
	a = PowOfFour_math(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 256
	a = PowOfFour_math(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 1024
	a = PowOfFour_math(cases)
	fmt.Println("case:", cases, "result:", a)
}
func TestPowOfFour_iter(t *testing.T) {
	cases := 0
	a := PowOfFour_iter(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 1
	a = PowOfFour_iter(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 4
	a = PowOfFour_iter(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 10
	a = PowOfFour_iter(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 256
	a = PowOfFour_iter(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 1024
	a = PowOfFour_iter(cases)
	fmt.Println("case:", cases, "result:", a)
}
