# Garbage Collection

[Understanding Java Garbage Collection](https://www.cubrid.org/blog/understanding-java-garbage-collection)

### Stop-The-World
JVM stop applicaiton to gc. Application will resume after gc did its job.




minor GC: delete young generation object
major GC: delete old generation object
permenant generation, call method area for class or interned strings





If old hold young, card table comes out.
card table is 512 byte chunk  with write barrier
This write barrier is a device that allows a faster performance for minor GC


### young generation
One Eden space
Two Survivor spaces

1. object survives many times of gc(move to other surviving space)
2. when survior space is full, survior space will be charged with a state for no data anymore
3. survive more times and move to old generation







### faster memeory allocation

    + bump-the-pointer
    + TLABs (Thread-Local Allocation Buffers)


    Bump-the-pointer technique tracks the last object allocated to the Eden space. That object will be located on `top of the Eden space`. And if there is an object created afterwards, `it checks only if the size` of the object is suitable for the Eden space. If the said object seems right, it will be placed in the Eden space, and the new object goes on top. So, when new objects are created, `only the lastly added object needs to be checked`, which allows much faster memory allocations. However, it is a different story if we consider a multithreaded environment. To save objects used by multiple threads in the Eden space for Thread-Safe, an inevitable lock will occur and the performance will drop due to the lock-contention. TLABs is the solution to this problem in HotSpot VM. This allows each thread to have a small portion of its Eden space that corresponds to its own share. As each thread can only access to their own TLAB, even the bump-the-pointer technique will allow memory allocations without a lock. 




### old generation
perform gc when data is full, varies by gc type







### GC types

    + Serial gc
        -XX:+UseSerialGC) , algo: mark-sweep-compact
        This GC type was created when there was only one CPU core on desktop computers. Using this serial GC will drop the application performance significantly. 
    + Parallel
        -XX:+UseParallelGC) , called `throughput GC` , as same as Serial GC but with multi-threads
    + Parallel Old GC (Parallel Compacting GC)
        -XX:+UseParallelOldGC) , algo: mark–summary–compact

    + Concurrent Mark & Sweep GC (or "CMS")
        -XX:+UseConcMarkSweepGC , called `low latency GC` , used when the response time from all applications is crucial. 
        short stop-the-world time, more memory and CPU, compaction step is not provided by default.

    + Garbage First(G1) GC