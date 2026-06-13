package tree

// LCA, Lowest Common Ancestor
// [Pattern]: [LCA] Find common ancestor between node A and node B
func LCA_BT(root, a, b *Node) *Node {

	if root == nil {
		return nil
	}

	matched := (root == a || root == b)

	ln := LCA_BT(root.Left, a, b)
	rn := LCA_BT(root.Right, a, b)

	if matched || (ln != nil && rn != nil) {
		return root
	}

	if ln != nil {
		return ln
	}

	if rn != nil {
		return rn
	}

	return nil
}

// [Variant]: [LCA] Find common ancestor in BST
func LCA_BST(root, p, q *Node) *Node {
	if p.Val > root.Val && q.Val > root.Val {
		return LCA_BST(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return LCA_BST(root.Left, p, q)
	}

	return root
}

// [Variant]: [LCA] Fidn the length between node A and node B.
func FindLengthBetweenTwoNodes(a *Node, b *Node) int {
	return 0
}
