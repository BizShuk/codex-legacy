package sort

import (
	"fmt"
	"testing"
)

func TestSort2Dcolumn(t *testing.T) {
	//[[],[4],[1],[0],[4,1],[4,0],[1,0],[4,1,0]]
	a1 := Sort2Dcolumn([][]int{
		{}, {4}, {1}, {0}, {1, 4}, {0, 4}, {0, 1}, {0, 1, 4},
	})
	fmt.Println(a1)
}
