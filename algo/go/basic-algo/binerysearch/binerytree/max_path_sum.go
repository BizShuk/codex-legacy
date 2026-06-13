package binerytree

import (
	"math"

	"github.com/bizshuk/algo/util"
)

// [Variant]: [Tree Lengh] Maximum sum path
// [Hint]: With negative number, 0 can be a special number that presenting not take anyone.
func MaxPathSum(root *Node) int {

	if root == nil {
		return 0
	}

	maxNum := root.Val

	var dfs func(root *Node) int

	dfs = func(root *Node) int {

		if root == nil {
			return math.MinInt
		}

		lv := dfs(root.Left)
		rv := dfs(root.Right)
		lm := util.Max(lv, 0)
		rm := util.Max(rv, 0)

		maxNum = util.Max(maxNum, root.Val+lm+rm)

		return root.Val + util.Max(lm, rm)

	}

	dfs(root)

	return maxNum
}
