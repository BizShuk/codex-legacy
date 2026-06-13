package golang

import "fmt"

type Var int

// [Pattern]: [Go Escape Variable] Escape Variable Principle
// if variable is generated inside method and expose outside of method => escape (cost: heap allogcate and gc)
// if the return value is calculated by pointer passed from outside, the compile? or runtime? might treat the return variable is in outside method stack scope. (no escape )
// Ref: https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d

// [Tip]: When a heap allocation occurs in a function being benchmarked, we’ll see the allocs/ops stat go up by one. It’s the job of the garbage collector to later free heap variables that are no longer referenced.

func NonEscape(v Var) {
	fmt.Println(v)
}

func Escape() *Var {
	// var x Var = 0
	// return &x
	return new(Var)
}
