package linklist

import (
	"fmt"
	"testing"
)

func TestDelete_Duplicates(t *testing.T) {
	case1 := Create_Simplelist([]int{1, 2, 3, 4, 5})
	case2 := Create_Simplelist([]int{1, 2, 2, 2, 5})
	case3 := Create_Simplelist([]int{1, 1, 1, 1, 1})
	case4 := Create_Simplelist([]int{})
	case5 := Create_Simplelist([]int{1, 2, 3, 4, 4})

	exec_Delete_Duplicates(case1)
	exec_Delete_Duplicates(case2)
	exec_Delete_Duplicates(case3)
	exec_Delete_Duplicates(case4)
	exec_Delete_Duplicates(case5)
}

func exec_Delete_Duplicates(cases *Simplelist) {
	head := cases.Head

	head = Delete_Duplicates(head)
	node := head
	for node != nil {
		fmt.Print(node.Val, "=>")
		node = node.Next
	}
	fmt.Println("nil")

}
