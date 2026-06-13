package sync

import (
	"time"
)

func TimeSleepAndCancelation() {

	var done bool

	go func() {
		for {
			println("tick")
			time.Sleep(1 * time.Second)
			if done {
				println("cancelled")
				return
			}
		}
	}()

	<-time.After(time.Second * 5)
	done = true
	<-time.After(time.Second * 5)
}
