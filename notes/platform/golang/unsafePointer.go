package golang

import (
	"fmt"
	"sync"
	"unsafe"
)

type myStruct struct {
	a int32
	b int32
}

func Unsafe_Pointer() {
	x := 5
	p := unsafe.Pointer(&x)
	fmt.Println(p)
}

func Low_level_Memory_Manipulation() {
	p1 := &myStruct{}
	p2 := (*int32)(unsafe.Pointer(p1))
	*p2 = 10
	fmt.Println(*p1)
}

func Performance_Optimization() {

	x1 := &[100]int{}
	x2 := (*int)(unsafe.Pointer(x1))
	*x2 = 10
	fmt.Println(*x2)
}

type MyStructA struct {
	a int
	b int
}

type MyStructB struct {
	c int
	d int
}

func ByPass_Type_Safety_Check() {

	var x MyStructA
	p := (*MyStructB)(unsafe.Pointer(&x))
	p.c = 10
	p.d = 20
	fmt.Println(x)
}

func Security_Vulnerability() {
	var data [1024]byte
	var size uintptr = 1024
	buffer := data[:size]
	p := (*[1 << 30]byte)(unsafe.Pointer(&buffer))[:size:size]
	fmt.Println(string(p))
}

func Undefined_Behavior() {
	// var x int
	// p := (*int)(unsafe.Pointer(uintptr(0x1234)))
	// *p = 10
	// fmt.Println(x)
}

func Maintenance() {
	var x myStruct
	p := (*int)(unsafe.Pointer(&x))
	*p = 10
	fmt.Println(x)
}

func increment(p *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*p++
}

func Concurrent_Issue() {
	var x int
	p := (*int)(unsafe.Pointer(&x))
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(p, &wg)
	}
	wg.Wait()
	fmt.Println(x)
}
