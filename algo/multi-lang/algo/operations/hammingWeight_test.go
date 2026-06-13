package op

import (
	"fmt"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	cases := 0
	a := HammingWeight(cases)
	fmt.Println("case:", cases, "result:", a)

	cases = 1
	a = HammingWeight(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 2
	a = HammingWeight(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 10
	a = HammingWeight(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 200
	a = HammingWeight(cases)
	fmt.Println("case:", cases, "result:", a)
}
