package op

import (
	"fmt"
	"testing"
)

func TestHappyNum(t *testing.T) {
	cases := 123
	a := HappyNum(cases)
	fmt.Println("case:", cases, "result", a)
	cases = 456
	a = HappyNum(cases)
	fmt.Println("case:", cases, "result", a)

	cases = 7
	a = HappyNum(cases)
	fmt.Println("case:", cases, "result", a)
	cases = 1
	a = HappyNum(cases)
	fmt.Println("case:", cases, "result", a)
	cases = 0
	a = HappyNum(cases)
	fmt.Println("case:", cases, "result", a)
	cases = 9
	a = HappyNum(cases)
	fmt.Println("case:", cases, "result", a)
}
func TestHappyNum_induct(t *testing.T) {
	cases := 123
	a := HappyNum_induct(cases)
	fmt.Println("case:", cases, "result", a)
	cases = 456
	a = HappyNum_induct(cases)
	fmt.Println("case:", cases, "result", a)

	cases = 7
	a = HappyNum_induct(cases)
	fmt.Println("case:", cases, "result", a)
	cases = 1
	a = HappyNum_induct(cases)
	fmt.Println("case:", cases, "result", a)
	cases = 0
	a = HappyNum_induct(cases)
	fmt.Println("case:", cases, "result", a)
	cases = 9
	a = HappyNum_induct(cases)
	fmt.Println("case:", cases, "result", a)
}
