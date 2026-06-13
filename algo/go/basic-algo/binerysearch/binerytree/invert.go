package binerytree

// [Variant]: [tree traverse] invert childrens
// [Tip]: tree root can be nil
func Invert(root *Node) *Node {
	if root == nil {
		return root
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	Invert(root.Left)
	Invert(root.Right)

	return root
}
