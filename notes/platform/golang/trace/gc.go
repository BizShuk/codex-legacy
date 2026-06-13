package trace

// [GODEBUG] GODEBUG=gctrace=1

// [Sample]
// gc9(2): 12+1+744+8 us, 2 -> 10 MB, 108615 (593983-485368) objects, 4825/3620/0 sweeps, 0(0) handoff, 6(91) steal, 16/1/0 yields
// gc10(2): 12+6769+767+3 us, 1 -> 1 MB, 4222 (593983-589761) objects, 4825/0/1898 sweeps, 0(0) handoff, 6(93) steal, 16/10/2 yields
// gc11(2): 799+3+2050+3 us, 1 -> 69 MB, 831819 (1484009-652190) objects, 4825/691/0 sweeps, 0(0) handoff, 5(105) steal, 16/1/0 yields

// [GC Data Pattern]:
// [gc number]([worker number]):
// gc9(2):
// [stop-the-world, sweeping],
// 12+1+744+8 us,
// [size after previous GC] -> [size before current GC] [Unit],
// 2 -> 10 MB
// [total number of objects in heap (including garbage)] ([total number of memory allocation]-[total number of free operations]) objects,
// 108615 (593983-485368) objects,
// [memory spans from previous GC]/[swept on demand or in background]/[swept during stop-the-world] sweeps
// 4825/3620/0 sweeps
// load balancing during parallel mark phase: there were 0 object handoff operations (0 objects were handoff), and 6 steal operations (91 objects were stolen)
// 0(0) handoff, 6(91) steal,
// effectiveness of parallel mark phase: there were total of 17 yield operations during waiting for another thread.
// 16/1/0 yields

// The GC is <mark-and-sweep> type. Total GC can be expressed as:

// [Tgc = Tseq + Tmark + Tsweep]
// [Tseq] // Tseq is time to stop user goroutines and some preparation activities (usually small)
// [Tmark = C1*Nlive + C2*MEMlive_ptr + C3*Nlive_ptr] // heap marking time, marking happens when all user goroutines are stopped, and thus can significantly affect latency of processing;
// [Nlive] is the number of live objects in the heap during GC, [MEMlive_ptr] is the amount of memory occupied by live objects with pointers, [Nlive_ptr] is the number of pointers in live objects.

// [Tsweep = C4*MEMtotal + C5*MEMgarbage] // heap sweeping time, sweeping generally happens concurrently with normal program execution, and so is not so critical for latency.
// [MEMtotal] is the total amount of heap memory, [MEMgarbage] is the amount of garbage in the heap.

// [GOGC], it will cause gc event when the heap grow value percentage, default 100(%)
// GOGC=50
// GOGC=200
// GOGC=1000
// GOGC=off
// 4MB -> 8MB -> GC ->
