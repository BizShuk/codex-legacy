package linklist

import (
	"errors"
	"fmt"

	"github.com/bizshuk/algo/linkedlist"
)

/*
type linklist interface {
	len() int
	head() *node
	tail() *node
	insert(node) error
	delete(node) bool
	is_tail() bool
	iterator() *node
}

type node interface {
	next() *node
	prev() *node
}
*/

type Simplelist struct {
	P    *linkedlist.Node
	Head *linkedlist.Node
	Tail *linkedlist.Node
	len  int
}

func Create_Simplelist(arr []int) *Simplelist {
	list := &Simplelist{}
	for i := 0; i < len(arr); i++ {
		list.Append(&linkedlist.Node{Val: arr[i]})
	}
	return list
}

func (l *Simplelist) error(s string) string {
	return s
}

func (l *Simplelist) Len() int {
	count := 0
	for p := l.Head; p != nil; p = p.Next {
		count++
	}
	return count
}

func (l *Simplelist) Iterator() *linkedlist.Node {
	l.P = l.P.Next
	return l.P
}

func (l *Simplelist) Printall() error {
	fmt.Printf("\nlist:")
	for i, c := 1, l.Head; c != nil; i, c = i+1, c.Next {
		fmt.Printf("%d:%s ,", i, c.Vals)
	}
	fmt.Printf("len:%d\n\n", l.len)
	return nil
}

func (l *Simplelist) Unshift(n *linkedlist.Node) int {
	if l.Head == nil {
		l.Tail = n
		l.P = n
	}

	n.Next = l.Head
	l.Head = n
	l.len++
	return l.len
}

func (l *Simplelist) Shift() (*linkedlist.Node, error) {
	if l.Head == nil {
		return nil, errors.New("No Element could be removed")
	}

	shifted := l.Head
	l.Head = shifted.Next
	l.len--
	return shifted, nil
}

func (l *Simplelist) Append(n *linkedlist.Node) int {
	l.len++
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		l.P = n
		return l.len
	}

	l.Tail.Next = n
	l.Tail = n
	return l.len
}

func (l *Simplelist) Remove() (*linkedlist.Node, error) {
	if l.Head == nil {
		return nil, errors.New("Mo Element could be removed")
	}

	c := l.Head
	if c == l.Tail {
		l.Head = nil
		l.Tail = nil
		return c, nil
	}

	for c.Next.Next != nil {
		c = c.Next
	}

	l.Tail = c
	c = c.Next
	l.Tail.Next = nil
	l.len--

	return c, nil
}

func (l *Simplelist) is_tail(n *linkedlist.Node) bool {
	if l.Tail == n {
		return true
	}
	return false
}

func (l *Simplelist) Reverse() error {

	if l.Head == l.Tail {
		return nil
	}

	next := l.Head.Next
	l.Tail = l.Head

	var pre *linkedlist.Node
	for next != nil {
		pre = l.Head
		l.Head = next
		next = l.Head.Next
		l.Head.Next = pre
	}
	l.Tail.Next = nil

	return nil
}

func (l *Simplelist) Copy() *Simplelist {
	lnode, newlist := l.Head, &Simplelist{}

	for lnode != nil {
		newnode := &linkedlist.Node{Val: lnode.Val}
		newlist.Append(newnode)
		lnode = lnode.Next
	}
	return newlist
}
