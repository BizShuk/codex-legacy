package sync

import "time"

func RaceCondition() {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			counter = counter + 1 // counter has race condition
		}()
	}

	time.Sleep(1 * time.Second)
	println(counter)
}
