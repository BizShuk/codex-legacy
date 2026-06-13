package channel

import "fmt"

// [Pattern]: [Go Channel]
// [Hint]: Set as buffered channel to prevent publiser hanging.
// [Hint]: Only sender should close channels, only close when receiver must be told
// [Hint]: Panic when close nil channel

// [Variant]: [Go Channel] Buffered Channel
func BufferedChannel() {
	ch := make(chan int, 100)
	defer close(ch)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// v, ok := <-ch // Simple fetch from channel
	for v := range ch { // Util ch gets closed
		fmt.Println(v)
	}
}

// [Variant]: [Go Channel] Unbuffered Channel
func UnbufferedChannel() {
	ch := make(chan int)
	done := make(chan int)

	defer close(ch)
	// v, ok := <-ch

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-done:
		return
	}
}
