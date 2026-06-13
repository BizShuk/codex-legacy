package summaryRange

import (
	"fmt"
	"testing"
)

func TestSummaryRange(t *testing.T) {
	a1 := summaryRanges([]int{0, 2, 3, 6, 8, 10, 11, 12, 13, 14, 15, 20})
	fmt.Println(a1)
}
