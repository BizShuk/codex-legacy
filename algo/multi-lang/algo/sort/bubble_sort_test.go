package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	case1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	case2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	case3 := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	case4 := []int{1, 6, 2, 7, 3, 8, 4, 9, 5, 10}

	case5 := []int{1, 2, 3, 4, 5, 10, 9, 8, 7, 6}
	case6 := []int{5, 4, 3, 2, 1, 10, 9, 8, 7, 6}

	case7 := []int{9, 1, 6, 4, 10, 8, 7, 2, 3, 5}
	case8 := []int{3, 4, 5, 9, 10, 7, 6, 8, 1, 2}
	case9 := []int{5, 10, 2, 3, 1, 8, 9, 6, 4, 7}
	var ct, et int
	ct, et = BubbleSort(case1)
	fmt.Println(case1, ct, et)

	ct, et = BubbleSort(case2)
	fmt.Println(case2, ct, et)

	ct, et = BubbleSort(case3)
	fmt.Println(case3, ct, et)

	ct, et = BubbleSort(case4)
	fmt.Println(case4, ct, et)

	ct, et = BubbleSort(case5)
	fmt.Println(case5, ct, et)

	ct, et = BubbleSort(case6)
	fmt.Println(case6, ct, et)

	ct, et = BubbleSort(case7)
	fmt.Println(case7, ct, et)

	ct, et = BubbleSort(case8)
	fmt.Println(case8, ct, et)

	ct, et = BubbleSort(case9)
	fmt.Println(case9, ct, et)

}
