package bit

import (
	"fmt"
	"testing"
)

func TestBitSequence(t *testing.T) {
	fmt.Println("")
	var a []int
	for i := 0; i < 30; i++ {
		a = countBits(i)
		fmt.Println("cases:", i, "result:", a)
	}
	cases := 123456
	a = BitSequence(cases)
	fmt.Println("cases:", cases, "result:", a)

	BitNumberAdd(1, 1)
	BitNumberAdd(3, 4)
	BitNumberAdd(1, 5)
	BitNumberAdd(2, 2)

	BitNumberAdd_rev(123, 321)
	BitNumberAdd_rev(1, 31)
	BitNumberAdd_rev(13, 21)
	BitNumberAdd_rev(3, 321)
}
