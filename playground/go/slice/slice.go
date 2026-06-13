package main

import (
	"fmt"
)

func main() {
	intlist := []int{1, 2, 3, 4, 5, 6, 6}

	intlist = append(intlist, 7)
	fmt.Printf("%+v %d\n", intlist, len(intlist))

	intlist = append([]int{3}, intlist...)
	fmt.Printf("%+v %d\n", intlist, len(intlist))

	for {
		if len(intlist) == 1 {
			break
		}
		intlist = intlist[1:]
		fmt.Printf("%+v %d\n", intlist, len(intlist))
	}
}
