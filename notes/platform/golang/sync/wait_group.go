package sync

import (
	"sync"
)

// [Pattern]: [Go] WaitGroupt
func WaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(num int) {
			Output(num)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// [Hint]: The value of i is changing, most of time, it'll be last value of for loop. But it can be between start and end value. Unless it is synchronous programe.
func WaitGroup_Bad() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			// Output(i) // [Issue]: Bad example
			wg.Done()
		}()
	}
	wg.Wait()
}

func Output(i int) {
	println("Output:", i)
}
