package stack

import "sort"

// [Pattern]: [Monotonic Stack] Car fleet reach end by group
// [Hint]: How to sort with other same-indexed array

type Pair struct {
	Remain   float64 // [Hint]: With accending Position, the remaining time can be used to check whether cars bump into another
	Position int
}

func CarFleet(target int, position []int, speed []int) int {
	pairs := make([]Pair, len(position))
	for i := range position {
		pairs[i] = Pair{float64(target-position[i]) / float64(speed[i]), position[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Position < pairs[j].Position
	})

	stack := []Pair{}

	for _, v := range pairs { // Alternative: use prevTime to comapre might be simpler
		for len(stack) > 0 {
			if stack[len(stack)-1].Remain <= v.Remain {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}

		stack = append(stack, v)
	}

	return len(stack)
}
