// what wrong
package tree

import (
	"fmt"

	"github.com/bizshuk/code_sandbox/go/list"
)

type Tree struct {
	root *TreeNode
}

func (t *Tree) GetHeight() {

}

func (t *Tree) Insert() {

}

func (t *Tree) Delete() {

}

// null = 0
// root = 1
func TreeHeight_byindex(n int) (h int) {
	h = 0
	if n <= 0 {
		return 0
	}
	for n/2 >= 1 {
		h++
		n /= 2
	}

	h++
	return h
}

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("%s", t.Val)
}

func (t *TreeNode) MaxDepth() int {
	if t == nil {
		return 0
	}

	left_max := 0
	if t.Left != nil {
		left_max = t.Left.MaxDepth()
	}
	right_max := 0
	if t.Right != nil {
		right_max = t.Right.MaxDepth()
	}

	if left_max > right_max {
		return left_max + 1
	}
	return right_max + 1
}

func (t *TreeNode) MinDepth() int {
	if t == nil {
		return 0
	}

	left_min := 0
	if t.Left != nil {
		left_min = t.Left.MinDepth()
	}

	right_min := 0
	if t.Right != nil {
		right_min = t.Right.MinDepth()
	}

	if left_min == 0 || right_min == 0 {
		return left_min + right_min + 1
	}

	if left_min < right_min {
		return left_min + 1
	}
	return right_min + 1
}

func IsSameTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val && IsSameTree(t1.Left, t2.Left) && IsSameTree(t1.Right, t2.Right)
}

func (t *TreeNode) HasPathSum(sum int) bool {
	if t.Left == nil && t.Right == nil {
		return sum == t.Val
	}

	return (root.Left != nil && root.Left.HasPathSum(sum-t.Val)) ||
		(root.Right != nil && root.Right.HasPathSum(sum-t.Val))

}

func (t *TreeNode) Show() {
	l := list.List{}
	l.Push_back_i(t)
	var tmp *TreeNode

	for {
		if l.IsEmpty() {
			break
		}
		tmp = l.DeQueue().Val.(*TreeNode)

		fmt.Println(tmp.Val, " => left:", tmp.Left, " , right:", tmp.Right)

		if tmp.Left != nil {
			l.Push_back_i(tmp.Left)
		}

		if tmp.Right != nil {
			l.Push_back_i(tmp.Right)
		}
	}
}
