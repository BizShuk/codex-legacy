package receiver

import (
	"fmt"
)

// TODO: clean up
type Sinterface interface {
	Get()
}

type PointerReceiverStructWithBracket struct {
	T string
}

func (p *PointerReceiverStructWithBracket) Get() {
	fmt.Println(p.T)
}

type ValueReceiverStructWithBracket struct {
	T string
}

func (v ValueReceiverStructWithBracket) Get() {
	fmt.Println(v.T)
}

type PointerReceiverStructWithoutBracket string

func (p *PointerReceiverStructWithoutBracket) Get() {
	fmt.Println(*p)
}

type ValueReceiverStructWithoutBracket string

func (v ValueReceiverStructWithoutBracket) Get() {
	fmt.Println(v)
}

func ReceiverTransfer() {
	// cannot use PointerReceiverStructWithBracket literal (type PointerReceiverStructWithBracket) as type Sinterface in argument to ReceiverInterpret:
	// PointerReceiverStructWithBracket does not implement Sinterface (Get method has pointer receiver)
	// ReceiverInterpret(PointerReceiverStructWithBracket{})

	ReceiverInterpret(&PointerReceiverStructWithBracket{})

	ReceiverInterpret(ValueReceiverStructWithBracket{})

	ReceiverInterpret(&ValueReceiverStructWithBracket{})

	// invalid composite literal type PointerReceiverStructWithoutBracket
	// ReceiverInterpret(PointerReceiverStructWithoutBracket{})

	// invalid composite literal type PointerReceiverStructWithoutBracket
	// ReceiverInterpret(&PointerReceiverStructWithoutBracket{})

	// invalid composite literal type ValueReceiverStructWithoutBracket
	// ReceiverInterpret(ValueReceiverStructWithoutBracket{})

	// invalid composite literal type ValueReceiverStructWithoutBracket
	// ReceiverInterpret(&ValueReceiverStructWithoutBracket{})

}

func ReceiverInterpret(s Sinterface) {
	s.Get()
}
