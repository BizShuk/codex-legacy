package trace

// [GODEBUG] GODEBUG=schedtrace=1000 // Period of output(ms)
// [GODEBUG] GODEBUG=scheddetail=1 // Prints detailed info about every goroutine, worker thread and processor

// Insights into dynamic behavior of the goroutine scheduler and allow to debug load balancing and scalability issues
// The scheduler trace is useful when a program does not scale linearly with GOMAXPROCS and/or does not consume 100% of CPU time

// [Sample]
// SCHED 1004ms: gomaxprocs=4 idleprocs=0 threads=11 idlethreads=4 runqueue=8 [0 1 0 3]
// SCHED 2005ms: gomaxprocs=4 idleprocs=0 threads=11 idlethreads=5 runqueue=6 [1 5 4 0]
// SCHED 3008ms: gomaxprocs=4 idleprocs=0 threads=11 idlethreads=4 runqueue=10 [2 2 2 1]

// [Scheduler Pattern]:
// SCHED 1004ms: [time since program start]
// gomaxprocs=4 [the current value of GOMAXPROCS]
// idleprocs=0 [the number of idling processors]
// threads=11 [the total number of worker threads created by the scheduler]
// threads can be in 3 states:
// 1. execute Go code (gomaxprocs-idleprocs)
// 2. execute syscalls/cgocalls
// 3. idle
// idlethreads=4 [the number of idling worker threads]
// runqueue=8 [the length of global queue with runnable goroutines]
// [0 1 0 3] [lengths of per-processor queues with runnable goroutines] Sum of lengths of global and local queues represents the total number of goroutines available for execution.
