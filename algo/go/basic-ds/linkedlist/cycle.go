package linkedlist

// [Variant]: [Cycle Detection] [Floyd's cycle-finding] findDuplicateNum with array
func HasCycle(head *Node) bool {
	s, f := head, head
	for f != nil && f.Next != nil {
		s, f = s.Next, f.Next.Next
		if s == f {
			return true
		}
	}
	return false
}

// [Variant]: [Disjoint Set Union]
