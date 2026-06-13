package pprof

// [Flag] --cpuprofile=<profile name>

// case:
// [runtime.mallocgc] => excessive amount of small memory allocations => check memory profiler
// [channel operations], [sync.Mutex] code and other synchronization primitives or [System component] => contention => eliminate frequently accessed shared resources. Common techniques for this include sharding/partitioning, local buffering/batching and <copy-on-write> technique.
// [syscall.Read/Write] =>  excessive amount of small reads and writes => Bufio wrappers around
// [GC component] =>  either allocates too many transient objects or heap size is very small => check Garbage Collector Tracer and Memory Profiler
