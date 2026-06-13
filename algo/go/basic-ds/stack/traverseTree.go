package stack

// [Question]: Generate combination "(())" from top-down
// [Hint]: tree left child '(', right child ')', and the tree path is depend on combinations.
func GenerateParenthesis2(n int) []string {
	return Dfs2([]byte{}, n, n)
}

// i: '(' count, j: ')' count
func Dfs2(b []byte, i, j int) []string {
	if i == 0 && j == 0 {
		return []string{string(b)}
	}
	var s1, s2 []string

	if i > 0 {
		b = append(b, '(')
		s1 = Dfs2(b, i-1, j)
		b = b[:len(b)-1]
	}

	if j > i {
		s2 = Dfs2(append(b, ')'), i, j-1)
	}
	return append(s1, s2...)
}

func GenerateParenthesis(n int) []string {
	stack := []string{}

	var dfs func(s string, i, j int)

	dfs = func(s string, i, j int) {
		if i == 0 && j == 0 {
			stack = append(stack, s)
		}

		if i > 0 {
			dfs(s+"(", i-1, j)
		}

		if j > i {
			dfs(s+")", i, j-1)
		}
	}

	dfs("", n, n)

	return stack
}
