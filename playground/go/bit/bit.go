package bit

import (
	"fmt"
)

func countBits(num int) []int {
	ret := []int{0}
	for i := 0; i <= num; i++ {
		if i == 0 {
			ret[i] = 0
			continue
		}
		ret = append(ret, ret[i>>1]+(i&1))
	}
	return ret

}

func BitSequence(num int) []int {
	var ret = make([]int, num+1)
	fmt.Println("")
	for i := 0; i <= num; i++ {
		ret[i] = ret[i>>1] + (i & 1)
	}
	return ret
}

func BitNumberAdd(n1 int, n2 int) int {

	sum := n1
	for n2 != 0 {
		sum = n1 ^ n2
		n2 = (n1 & n2) << 1
		n1 = sum
	}

	return sum
}

func BitNumberAdd_rev(n1 int, n2 int) int {
	if n2 == 0 {
		return n1
	}
	return BitNumberAdd_rev(n1^n2, (n1&n2)<<1)
}
