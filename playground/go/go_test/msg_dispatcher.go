package main

import (
	"fmt"
	"time"
)

func say(worker string) {
	for i := 1; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(worker, ":", i)
	}
}
func main() {
	fmt.Println("test")

	var t = []string{"a", "v", "f", "e", "k"}
	for i := range t {
		//	go say(i)
		fmt.Println(i)
	}
	say("l")
}
