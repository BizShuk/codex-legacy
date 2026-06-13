package btree

type Node struct {
	Value interface{}
	left  *Node
	right *Node
}

type Btree struct {
	root *Node
}

func (n *Node) Is_leaf() bool {
	if n.left == nil && n.right == nil {
		return true
	}
	return false
}
