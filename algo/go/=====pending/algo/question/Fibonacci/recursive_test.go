package Fibonacci

import (
	"fmt"
	"log"
	"testing"
)

func TestFibonaccirecursive(t *testing.T) {
	log.Println("Fibonacci recursive:")

	var n, goal int

	goal = 20
	n = Fibonacci_recursive(goal)
	fmt.Printf("n=%d:%d\n", goal, n)
	if n != 6765 {
		t.Error("Fibonacci_recursive:20 failed")
	}

	goal = 30
	n = Fibonacci_recursive(goal)
	fmt.Printf("n=%d:%d\n", goal, n)
	if n != 832040 {
		t.Error("Fibonacci_recursive:30 failed")
	}

	goal = 40
	n = Fibonacci_recursive(goal)
	fmt.Printf("n=%d:%d\n", goal, n)
	if n != 102334155 {
		t.Error("Fibonacci_recursive:40 failed")
	}

}
