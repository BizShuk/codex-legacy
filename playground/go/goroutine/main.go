package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func handler(index int, t chan int) {
	// with sequential
	for t_request := range t {
		fmt.Println(index, t_request)
		wg.Done()
	}
	// without sequential
	/*
		for {
			t_request := <-t
			wg.Done()
			fmt.Println(index, t_request)
		}
	*/
}

func send_request(c chan int) {

	t := make(chan int, 100)

	for i := 0; i < 10; i++ {
		go handler(i, t)
	}
	for i := 0; i <= 1000; i++ {
		wg.Add(1)
		t <- i
	}

	wg.Wait()

	close(t)
	c <- 1
}
func main() {
	c := make(chan int)
	go send_request(c)
	<-c

	close(c)
}
