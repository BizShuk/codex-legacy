package sync

import (
	"sync"
)

// [Pattern]: [Go] sync.Once, function block only executes once

func SyncOnce() {

	var once sync.Once
	onceBody := func() {
		println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
