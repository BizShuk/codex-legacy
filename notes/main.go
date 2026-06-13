package main

import "fmt"

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	// sort.SortInts()	// golang/fundation/sort/int.go

	// golang/fundation
	// fundation.LoopEndValueWithPointerReceiver()
	// fundation.InitOrderOnVariables()
	// fundation.BasicSlice()
	// fundation.FmtInterface()
	// receiver.TypeAddress()
	// receiver.ReceiverInterface()
	// fundation.NumberPresent()

	// x := 0xff
	// fmt.Printf("%064b", x)

	// fundation.CharSlice()
	// sync.SyncCond()

	m := map[int]bool{-1: true}
	fmt.Println(len(m))
	for k, v := range m {

		m[k+1] = true
		fmt.Println(k, v)
	}

}
