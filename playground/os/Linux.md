# Linux

### page cache
ram, 讀寫檔案時 特別是 binary lib都會用這個, 當有dirty file要flush到disk時也會用這個

### Control Groups, cgroup
Reserved resources of each thread and processes control by cgroup
    For example, each thread will take between 320K (32 bit JVM) and 1024K (64 bit JVM) memory for its stack



### Memory Management
he kernel swap daemon (kswapd).


##### Kernel Swap Daemon

    When run out of physical memory, in order to free memory.
    `Kernel Swap Daemon (kswapd)` , a kernel thread, no virtual memory, started by init process
    + swap memeory
    + keep the memory management system operating efficiently.

    free_pages_high and free_pages_low to determine if it should free memory (set at system startup time). If the number of free pages fall into these two variables, kswapd will do following 3 things.
        - Reducing the size of the buffer and page caches,
        - Swapping out System V shared memory pages,
        - Swapping out and discarding pages.

    if lower than free_pages_low, free 6 pages before next run and 3 pages in turns. sleep when job is done, sleep half its usual time when it's still lower than free_pages_low.


### IPC, Inter Process Communication

    - shared memory
    - message passing (throught kernel)
        + direct, to specific pid
        + indirect, use mailbox or a queue
        Blocking send and blocking receive
        Non-blocking send and Non-blocking receive
        Non-blocking send and Blocking receive (Mostly used)

        Posix : uses shared memory method.
        Mach : uses message passing
        Windows XP : uses message passing using local procedural calls


        There are various mechanism:
        - Pipe
        - Socket
        - Remote Procedural calls (RPCs)

