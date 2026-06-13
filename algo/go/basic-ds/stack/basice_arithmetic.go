package stack

import "strconv"

// [Pattern]: [First Class Function] Evaluate arithmetic result from string

// [Hint]: 1. Leverage existing space for tmp result.
// [Hint]: 2. Clear what we use string, byte, rune, int, uint8

func add(x, y int) int      { return x + y }
func minus(x, y int) int    { return x - y }
func multiply(x, y int) int { return x * y }
func divide(x, y int) int   { return x / y }

var Ops = map[string]func(x, y int) int{
	"+": add,
	"-": minus,
	"*": multiply,
	"/": divide,
}

func EvalRPN_general(tokens []string) int {
	stack := []int{}
	for _, s := range tokens {
		if arithFunc, ok := Ops[s]; ok {
			stack[len(stack)-2] = arithFunc(stack[len(stack)-2], stack[len(stack)-1])
			stack[len(stack)-1] = 0
			stack = stack[:len(stack)-1]
			continue
		}

		num, _ := strconv.Atoi(s)
		stack = append(stack, num)

	}
	return stack[0]
}

func EvalRPN_switch(tokens []string) int {
	stack := []int{}
	for _, s := range tokens {

		switch s {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
		default: // push stack
			num, _ := strconv.Atoi(s)
			stack = append(stack, num)
			continue
		}

		stack[len(stack)-1] = 0
		stack = stack[:len(stack)-1]
	}
	return stack[0]
}
