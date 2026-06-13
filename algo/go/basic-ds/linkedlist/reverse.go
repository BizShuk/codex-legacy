package linkedlist

// [Pattern]: [Reverse LinkedList] Reverse LinkedList
func Reverse(head *Node) *Node { // Return: tail node
	var prev *Node

	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}
