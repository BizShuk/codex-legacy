package main

import (
	"fmt"
	"strconv"
	"time"
)

var tx chan string

func main() {
	tx = make(chan string, 10)

	w1 := writeClosure("w1")
	w2 := writeClosure("w2")
	w3 := writeClosure("w3")
	go w1()
	go w2()
	go w3()

	time.Sleep(1 * time.Second)

	r1 := readClosure("r1")
	go r1()
	r2 := readClosure("r2")
	go r2()

	time.Sleep(1 * time.Second)
}

func writeClosure(id string) func() {
	var count int = 0
	return func() {
		for {
			// skipping if capacity is full
			// select {
			// case tx <- id + ":" + strconv.Itoa(count):
			// 	fmt.Printf("Send by %v with value %v\n", id, count)
			// default:
			// 	fmt.Printf("Send by %v skipping value %v\n", id, count)
			// }

			// blocking if capacity is full
			tx <- id + ":" + strconv.Itoa(count)
			fmt.Printf("Send by %v with value %v\n", id, count)

			count++
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func readClosure(id string) func() {
	return func() {
		for x := range tx {
			fmt.Printf("Receive by %v with value %v\n", id, x)
		}
	}
}
