// test godoc
package linklist

import (
	"fmt"
	"testing"

	"github.com/bizshuk/algo/linkedlist"
)

func TestSimple(t *testing.T) {
	n1 := &linkedlist.Node{Vals: "t1"}
	n2 := &linkedlist.Node{Vals: "t2"}
	n3 := &linkedlist.Node{Vals: "t3"}
	n4 := &linkedlist.Node{Vals: "t4"}
	n5 := &linkedlist.Node{Vals: "t5"}
	n6 := &linkedlist.Node{Vals: "t6"}
	n7 := &linkedlist.Node{Vals: "t7"}

	simplelist := Simplelist{}
	simplelist.Printall()

	sn1, err := simplelist.Remove()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%v\n", sn1)
	simplelist.Printall()

	sn2, err := simplelist.Shift()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%v\n", sn2)
	simplelist.Printall()

	len := simplelist.Unshift(n7)
	simplelist.Printall()

	fmt.Printf("%v\n", simplelist.Append(n1))
	fmt.Printf("%v\n", simplelist.Append(n2))
	fmt.Printf("%v\n", simplelist.Append(n3))
	fmt.Printf("%v\n", simplelist.Append(n4))
	fmt.Printf("%v\n", simplelist.Append(n5))
	fmt.Printf("list len:%v\n", simplelist.Len())
	simplelist.Printall()

	i := simplelist.Head
	for i != nil {
		fmt.Printf("%s\n", i.Vals)
		i = simplelist.Iterator()
	}

	simplelist.Printall()

	sn3, err := simplelist.Shift()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%v\n", sn3)
	simplelist.Printall()

	len = simplelist.Unshift(n6)
	fmt.Printf("%d\n", len)
	simplelist.Printall()

	sn4, err := simplelist.Remove()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%v\n", sn4)
	simplelist.Printall()

	fmt.Println("Linklist reverse")
	_ = simplelist.Reverse()
	simplelist.Printall()

	_ = simplelist.Reverse()
	simplelist.Printall()

	fmt.Println("Linklist copy")
	newcopy := simplelist.Copy()
	_ = newcopy.Printall()
	newcopy.P.Vals = "99"
	simplelist.Printall()
	_ = newcopy.Printall()

}
