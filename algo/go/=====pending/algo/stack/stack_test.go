package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := &stack{}

	s.Push(3)
	s.Push(4)
	s.Push(5)

	for i := 1; i < 3; i++ {
		err, popd := s.Pop()
		fmt.Printf("%d %v\n", popd, err)
	}

}
