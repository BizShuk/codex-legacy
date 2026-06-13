package stack

import "github.com/bizshuk/algo/util"

// [Pattern]: [Min Stack] keep one value along with current stack

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(val int) {
	min := val
	if len(this.stack) > 0 { // [Condition]: length == 0
		min = util.Min(this.minStack[len(this.minStack)-1], val)
	}

	this.stack = append(this.stack, val)
	this.minStack = append(this.minStack, min)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int { // [Condition]: length == 0
	return this.minStack[len(this.minStack)-1]
}
