package backtracking

import (
	"fmt"
	"testing"
)

func TestIsValidParenthese(t *testing.T) {
	case1 := "(){}[]"
	case2 := "("
	case3 := "[()]"
	case4 := "([)]"
	case5 := "({[]}){}[]"

	exec_IsValidParenthese(case1)
	exec_IsValidParenthese(case2)
	exec_IsValidParenthese(case3)
	exec_IsValidParenthese(case4)
	exec_IsValidParenthese(case5)

}

func exec_IsValidParenthese(cases string) {
	a := IsValid_Parenthese(cases)
	fmt.Println("Case:", cases)
	fmt.Println("Result:", a)
}
