# Internet

- C10M: Defending the Internet at scale by `Robert David Graham`

### TCP vs UDP
- connection vs connectionless
- order: ordered vs non-ordered
- Reliability: good vs bad
- Data boundary: byte stream no boundary vs boundary with package
- Speed: slow vs fast
- Weight: heavy vs light
- Header size: 20bytes : 

### socket server
[Java Socket server sample](http://tutorials.jenkov.com/java-multithreaded-servers/index.html)

##### Thread Pool type
    + Single Thread Executor : A thread pool with only one thread. So all the submitted tasks will be executed sequentially. Method : Executors.newSingleThreadExecutor()

    + Cached Thread Pool : A thread pool that creates as many threads it needs to execute the task in parrallel. The old available threads will be reused for the new tasks. If a thread is not used during 60 seconds, it will be terminated and removed from the pool. Method : Executors.newCachedThreadPool()

    + Fixed Thread Pool : A thread pool with a fixed number of threads. If a thread is not available for the task, the task is put in queue waiting for an other task to ends. Method : Executors.newFixedThreadPool()

    + Scheduled Thread Pool : A thread pool made to schedule future task. Method : Executors.newScheduledThreadPool()

    + Single Thread Scheduled Pool : A thread pool with only one thread to schedule future task. Method : Executors.newSingleThreadScheduledExecutor()


##### Nonblocking
- partial message
- We want to copy message data around as little as possible. The more copying, the lower performance.
- We want full messages to be stored in consecutive byte sequences to make parsing messages easier.

##### Non Blocking on regular files
    [ref](https://www.remlab.net/op/nonblock.shtml)
    Reading from a regular file might take a long time. For instance, if it is located on a busy disk, the I/O scheduler might take so much time that the user will notice that the application is frozen.

    Nevertheless, non-blocking mode will not fix it. It will simply not work. Checking a file for readability or writeability always succeeds immediately. If the system needs time to perform the I/O operation, it will put the task in non-interruptible sleep from the read or write system call. In other words, if you can assume that a file descriptor refers to a regular file, do not waste your time (or worse, other people's time) in implementing non-blocking I/O.

    The only safe way to read data from or write data to a regular file while not blocking a task... consists of not performing the operation - not in that particular task anyway. Concretely, you need to create a separate thread (or process), or use asynchronous I/O (functions whose name starts with aio_). Whether you like it or not, and even if you think multiple threads suck, there are no other options.

    An alternative, of course, involves reading small chunks of data at once, and handling other events in-between. Then again, even reading a single byte can take a long time, if said byte was not read ahead by the operating system.
