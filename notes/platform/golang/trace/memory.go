package trace

// [GODEBUG] GODEBUG=allocfreetrace=1

// [Memory Statistic] https://github.com/golang/go/wiki/performance#memory-statistics
// https://pkg.go.dev/runtime#MemStats

// HeapAlloc - current heap size.
// HeapSys - total heap size.
// HeapObjects - total number of objects in the heap.
// HeapReleased - amount of memory released to the OS; runtime releases to the OS memory unused for 5 minutes, you can force this process with runtime/debug.FreeOSMemory.
// Sys - total amount of memory allocated from OS.
// Sys-HeapReleased - effective memory consumption of the program.
// StackSys - memory consumed for goroutine stacks (note that some stacks are allocated from heap and are not accounted here, unfortunately there is no way to get total size of stacks (https://code.google.com/p/go/issues/detail?id=7468)).
// MSpanSys/MCacheSys/BuckHashSys/GCSys/OtherSys - amount of memory allocated by runtime for various auxiliary purposes; they are generally not interesting, unless they are too high.
// PauseNs - durations of last garbage collections.

// [Heap Dump] https://github.com/golang/go/wiki/performance#heap-dumper
