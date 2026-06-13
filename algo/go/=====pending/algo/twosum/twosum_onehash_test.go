package twosum

import (
	"testing"
)

func TestTwosumhash(t *testing.T) {
	paired := twoSum_hash([]int{3, 2, 4}, 6)
	answer := []int{1, 2}
	if !compare_slice(paired, answer) {
		t.Error(paired, answer, "Not matched")
	}

	paired = twoSum_hash([]int{0, 2, 4, 0}, 0)
	answer = []int{0, 3}
	if !compare_slice(paired, answer) {
		t.Error(paired, answer, "Not matched")
	}
}

func compare_slice(a, b []int) bool {
	bmap := make(map[int]int)

	for _, v := range b {
		bmap[v] = 1
	}

	for _, v := range a {
		if _, ok := bmap[v]; !ok {
			return false
		}
	}
	return true
}
