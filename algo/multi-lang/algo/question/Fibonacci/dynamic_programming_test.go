package Fibonacci

import (
	"fmt"
	"log"
	"testing"
)

func TestFibonaccidynamicprogramming(t *testing.T) {
	var n, goal int
	log.Println("Fibonacci dynamic programming:")
	goal = 20
	n = Fibonacci_dynamicprogramming(goal)
	fmt.Printf("n=%d:%d\n", goal, n)
	if n != 6765 {
		t.Error("Fibonacci_dynamicprogramming:20 failed")
	}

	goal = 30
	n = Fibonacci_dynamicprogramming(goal)
	fmt.Printf("n=%d:%d\n", goal, n)
	if n != 832040 {
		t.Error("Fibonacci_dynamicprogramming:30 failed")
	}

	goal = 40
	n = Fibonacci_dynamicprogramming(goal)
	fmt.Printf("n=%d:%d\n", goal, n)
	if n != 102334155 {
		t.Error("Fibonacci_recursive:40 failed")
	}

}
