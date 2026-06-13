package sync

import (
	"sync"
)

// [Pattern]: [Go] Mutex Basic
func MutexBasic() {
	counter := 0
	var mu sync.Mutex

	// [Solution] 1: Sequential Transactions
	for i := 0; i < 5; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock() // [Tip]: OK to use defer unlock here
			counter = counter + 1
			println(counter)
		}()
	}
	// [Solution] 2: Parallel Transactions
	go func() {
		for i := 0; i < 5; i++ {
			mu.Lock()
			// defer mu.Unlock()	// [Tip]: Can't use defer unlock here
			counter = counter + 1
			println(counter)
			mu.Unlock()
		}
	}()
	// [Solution] 1: Waiting result, time.After channel
	// <-time.After(time.Second)

	// [Solution] 2: Waiting result, forloop time.Since
	// start := time.Now()
	// for time.Since(start) < time.Second {
	// }

	// [Solution] 3: Waiting result, WaitGroup
}
