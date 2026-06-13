package main

import "fmt"

var tx chan string

func main() {
	tx = make(chan string, 3)
	id := "w1"
	tx <- id + ":" + "3"

	for x := range tx {

		fmt.Println(x)
		if len(tx) == 0 {
			break
		}
	}

}
