package linkedlist

// [Pattern]: [Clone LinkedList] Clone linked list with random list
func CopyRandomList(head *Node) *Node {
	mapping := map[*Node]*Node{}

	preNewHead := &Node{} // point to new copied head
	var prev, curr, nCurr *Node = preNewHead, head, nil
	for curr != nil { // create a new copy list with Next pointer only
		prev.Next = &Node{Val: curr.Val}
		mapping[curr] = prev.Next
		prev, curr = prev.Next, curr.Next
	}

	curr, nCurr = head, preNewHead.Next

	for curr != nil { // mapping the random pointer to the new copied node
		nCurr.Random = mapping[curr.Random]
		curr, nCurr = curr.Next, nCurr.Next
	}

	return preNewHead.Next
}

// [Variant]: [Clone LinkedList] In-Place space
// 1. Chain new node after curr old node
// 2. Find new random node by next node of random node of old node
