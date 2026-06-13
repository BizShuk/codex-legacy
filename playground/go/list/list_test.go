package list

import (
	"fmt"
	"testing"
)

func TestListCopy(t *testing.T) {

	// list 1
	fmt.Println("List 1: 1 -> 2 -> 3 -> 4 -> 5")
	a1 := List{}
	a1.Push_back(&Listnode{Val: 1})
	a1.Push_back(&Listnode{Val: 2})
	a1.Push_back(&Listnode{Val: 3})
	a1.Push_back(&Listnode{Val: 4})
	a1.Push_back(&Listnode{Val: 5})
	a1.Show()

	// list 2
	fmt.Println("List 2: nil")
	a2 := List{}
	a2.Show()

	// list 3 , copy from a1
	a3 := a1.Copy()
	fmt.Println("List 1 with modification: 2 -> 2 -> 3 -> 4 -> 5")
	a1.Head.Val = a1.Head.Val.(int) + 1
	a1.Show()
	fmt.Println("List 3 ( copy from 1 ): 1 -> 2 -> 3 -> 4 -> 5")
	a3.Show()

	// list 4 , edited from a1 for arbitrary ptr
	fmt.Println("List 4: 2(3) -> 2(5) -> 3() -> 4(2) -> 5()")
	n1 := a1.Get(0)
	n2 := a1.Get(1)
	n3 := a1.Get(2)
	n4 := a1.Get(3)
	n5 := a1.Get(4)

	n1.Aptr = n3
	n2.Aptr = n5
	n4.Aptr = n2
	a1.Show()

	// list 5 , copy from list 4
	fmt.Println("List 5: 2(3) -> 2(5) -> 3() -> 4(2) -> 5() ")
	a5 := a1.CopyArbitrary()
	a5.Show()

	// addTwoNumber
	lt2 := List{}
	lt2.Push_back(&Listnode{Val: 2})
	lt2.Push_back(&Listnode{Val: 4})
	lt2.Push_back(&Listnode{Val: 3})

	lt3 := List{}
	lt3.Push_back(&Listnode{Val: 5})
	lt3.Push_back(&Listnode{Val: 6})
	lt3.Push_back(&Listnode{Val: 4})

	nt2 := lt2.Cur
	nt3 := lt3.Cur
	fmt.Println("addTwoNumber: nt2:", nt2.Val, ",nt3:", nt3.Val)

	nt_new := addTwoNumber(nt2, nt3, 0)
	nt_new.Show()

	// remove Nth from end
	nrn := make([]interface{}, 5)
	for i := 0; i < 5; i++ {
		nrn[i] = i
	}

	nrn1 := CreateList(nrn)
	nrn1.Show()
	RemoveNthFromEnd(nrn1, 1)
	nrn1.Show()
	nrn1 = RemoveNthFromEnd(nrn1, 4)
	nrn1.Show()
	nrn1 = RemoveNthFromEnd(nrn1, 1)
	nrn1.Show()
	nrn1 = RemoveNthFromEnd(nrn1, 1)
	nrn1.Show()
	nrn1 = RemoveNthFromEnd(nrn1, 1)
	nrn1.Show()

	a := &Listnode{Val: 1}

	atb := func(b *Listnode) {
		b = &Listnode{Val: 2}
		fmt.Println(b)

	}
	fmt.Println(a)
	atb(a)
	fmt.Println(a)

}
