package tree

import (
	"fmt"
	"testing"
)

func TestTreeHeight_byindex(t *testing.T) {
	cases := 0
	a := TreeHeight_byindex(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 1
	a = TreeHeight_byindex(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 2
	a = TreeHeight_byindex(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 25
	a = TreeHeight_byindex(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 4
	a = TreeHeight_byindex(cases)
	fmt.Println("case:", cases, "result:", a)

	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}

	n1.Left = n2
	n2.Left = n3
	n3.Right = n4
	n1.Right = n5
	n1.Show()
	fmt.Println("n1 max height:", n1.MaxDepth())
	fmt.Println("n1 min height:", n1.MinDepth())

}
