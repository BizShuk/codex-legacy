package fundation

import (
	"fmt"
	"sync"
)

type LockStruct struct {
	lock sync.Mutex
}

// [Pattern]: [Go] Mutex Sample
func UseMutex() {
	var t LockStruct // declaration has created a new instance
	t.lock.Lock()
	fmt.Println("Critical Session running...")
	t.lock.Unlock()
}
