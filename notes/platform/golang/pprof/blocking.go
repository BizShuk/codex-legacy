package pprof

// [Problem]:
// 1. No scale due to lack of work
// 2. Excessive blocking goroutine consumes CPU time

// [Solution]:
// 1. Use buffer
// 2. RWMutex instead of Mutex
// 3. Remove Mutext and use Copy-on-Write tech
// 4. Partition, one mutex for 1000 -> 10 mutext for 1000
// 5. Local cache and batch update
// 6. Sync.Pool

// TODO: document (sync.Pool sync map, sync cond)
// https://geektutu.com/post/hpg-sync-pool.html
