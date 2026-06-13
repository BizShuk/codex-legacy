package sync

import (
	"math/rand"
	"sync"
	"time"
)

// [Pattern]: [Go] sync.Condition
// Four main functions
// - NewCond(l Locker) *Cond
// - Broadcast(), Wake up all goroutine waiting for this locker
// - Signal(), Wake up only one of goroutines waiting for this locker
// - Wait(), Release one lock to process other one, should put in for loop with value check

func SyncCond() {
	rand.Seed(time.Now().UnixNano()) // [Noitce]: [Go] time rand.Seed
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	count, finished := 0, 0

	for i := 0; i < 20; i++ {
		go func() {
			vote := requestVote()
			cond.L.Lock()
			defer cond.L.Unlock()

			if vote {
				count++
			}
			finished++

			// Wake up cond.Wait to check
			cond.Signal() // cond.Broadcast(), both work since only one cond.Wait()
		}()
	}

	cond.L.Lock()
	for count < 5 && finished != 10 {
		cond.Wait()
		println("Condition:", count, finished)
	}
	cond.L.Unlock()

	if count >= 5 {
		println("received 5+ votes!")
	} else {
		println("lost")
	}
}

func requestVote() bool {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return rand.Int()%2 == 0
}
