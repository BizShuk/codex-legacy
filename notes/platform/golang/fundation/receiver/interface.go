package receiver

import "fmt"

func ReceiverInterface() {
	var x1 IFunc = F1{1} // value receiver interface
	x1.Update(2)
	x1.Call()

	// [Pattern]: [Go Pointer Receiver interface] No value type to pointer receiver interface
	var x2 IFunc = &F2{1} // pointer receiver interface
	x2.Update(2)
	x2.Call()
}

type IFunc interface {
	Update(x int)
	Call()
}

type F1 struct {
	val int
}

func (f F1) Update(x int) {
	// f.val = x // [Pattern]: [Go Receiver] No update to value receiver properties
}

func (f F1) Call() {
	fmt.Println(f.val)
}

type F2 struct {
	val int
}

func (f *F2) Update(x int) {
	f.val = x
}
func (f *F2) Call() {
	fmt.Println(f.val)
}
