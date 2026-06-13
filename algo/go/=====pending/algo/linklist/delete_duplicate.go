package linklist

import (
	"github.com/bizshuk/algo/linkedlist"
)

func Delete_Duplicates(head *linkedlist.Node) *linkedlist.Node {
	node := head
	if node == nil {
		return nil
	}

	for node != nil {
		if node.Next != nil && node.Val == node.Next.Val {
			node.Next = node.Next.Next
			continue
		}

		node = node.Next
	}
	return head
}
