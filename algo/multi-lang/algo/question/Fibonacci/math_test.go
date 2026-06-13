package Fibonacci

import (
	"fmt"
	"log"
	"testing"
)

func TestFibonaccimath(t *testing.T) {
	var n, goal float64
	log.Println("Fibonacci math:")
	goal = 20.0
	n = Fibonacci_math(goal)
	fmt.Printf("n=%f:%f\n", goal, n)
	if n != 6765.0 {
		t.Error("Fibonacci_math:20 failed")
	}

	goal = 30.0
	n = Fibonacci_math(goal)
	fmt.Printf("n=%f:%f\n", goal, n)
	if n != 832040.0 {
		t.Error("Fibonacci_math:30 failed")
	}

	goal = 40.0
	n = Fibonacci_math(goal)
	fmt.Printf("n=%f:%f\n", goal, n)
	if n != 102334155.0 {
		t.Error("Fibonacci_recursive:40 failed")
	}

}
