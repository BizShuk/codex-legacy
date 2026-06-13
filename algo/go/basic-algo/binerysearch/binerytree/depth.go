package binerytree

import "github.com/bizshuk/algo/util"

// [Pattern]: [tree depth] Find max depth of tree
// [Tip]: tree node could be nil
func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return util.Max(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}
