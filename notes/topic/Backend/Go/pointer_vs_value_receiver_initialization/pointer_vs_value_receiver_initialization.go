package main

import (
	"fmt"
)

type Sinterface interface {
	Get()
}

type pointerReceiverStructWithBracket struct {
	T string
}

func (p *pointerReceiverStructWithBracket) Get() {
	fmt.Println(p.T)
}

type valueReceiverStructWithBracket struct {
	T string
}

func (v valueReceiverStructWithBracket) Get() {
	fmt.Println(v.T)
}

type pointerReceiverStructWithoutBracket string

func (p *pointerReceiverStructWithoutBracket) Get() {
	fmt.Println(*p)
}

type valueReceiverStructWithoutBracket string

func (v valueReceiverStructWithoutBracket) Get() {
	fmt.Println(v)
}

func main() {
	// cannot use pointerReceiverStructWithBracket literal (type pointerReceiverStructWithBracket) as type Sinterface in argument to interpret:
	// pointerReceiverStructWithBracket does not implement Sinterface (Get method has pointer receiver)
	// interpret(pointerReceiverStructWithBracket{})

	interpret(&pointerReceiverStructWithBracket{})

	interpret(valueReceiverStructWithBracket{})

	interpret(&valueReceiverStructWithBracket{})

	// invalid composite literal type pointerReceiverStructWithoutBracket
	// interpret(pointerReceiverStructWithoutBracket{})

	// invalid composite literal type pointerReceiverStructWithoutBracket
	// interpret(&pointerReceiverStructWithoutBracket{})

	// invalid composite literal type valueReceiverStructWithoutBracket
	// interpret(valueReceiverStructWithoutBracket{})

	// invalid composite literal type valueReceiverStructWithoutBracket
	// interpret(&valueReceiverStructWithoutBracket{})

}

func interpret(s Sinterface) {
	s.Get()
}
