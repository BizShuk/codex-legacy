package golang

// TODO: [Go] goroutine調度
// TODO: [Go] memory modal
// ref:https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d

// [Pattern]: [Go memory] stack
// thread stack is typically fixed, e.g. a default of 8MB in many Linux environments.
// stack grow from high memory address g.stack.hi to low memory address g.stack.lo
// if function exceeds the side of stack, runtime will copy the stack to larger size of memory

// stack guard
// stack pointer
// frame pointer

// [Heap]

// [memory debug]
// cmd trace: https://pkg.go.dev/cmd/trace
// go test -run TestCopyIt -trace=copy_trace.out
// go tool trace copy_trace.out
// go tool trace -pprof=TYPE trace.out > TYPE.pprof
// - net: network blocking profile
// - sync: synchronization blocking profile
// - syscall: syscall blocking profile
// - sched: scheduler latency profile
// go tool pprof TYPE.pprof

// [GC]
// 1. Package-level allocated memory variables will never be collected.
// 2. Goroutine stack collected automatic (not by garbage collector) while exiting.
// 3. The garbage collector (GC) collects only allocated memory in a heap only if is not referenced i.e. unused memory blocks.

// [Variable Escape]
// Pointers or values as an argument to the function.
// Sending pointers or values containing pointers to channels: AS compiler cannot determine when goroutine will receive the data on a channel and when can be free.
// Pointers in a slice: slice may be on the stack but the referenced data escapes to the heap.
// Arrays of slices: It may get reallocated because an append can exceed its capacity.
// Methods on an interface type: As the interface gets type and value at runtime.
