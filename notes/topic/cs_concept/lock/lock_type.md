# Lock Types

[Deadlock]: Two processes/threads try to acquire lock from each other and not release lock in head.

[Livelock]: Like Spinlock, but without ending.

[Starvation]: Process/Thread never get lock.

[Race Condition]: Two processes/threads get the data and modify it at the same time. EX: counter issue.

[Spinlock]: keep cpu busy until acquiring lock

[Lock]: Lock in one process.

[Mutex]: Lock can be shared by multiple processes.

[Semaphore]: Similiar to Mutex but it can allow number of process/thread enter critical session

[Atomic OP]:

- Compare and Swap
- Fetch and Add
- Test and Set

[Algo if lock free]:

- Dekker
- Peterson
