package backtracking

// generate number of parenthese combination
// () for 1
// (()), ()() for 2
// ((())), (()()) , (())() , ()(()) , ()()() for 3
// ...
// ...
func Parenthese_Combine(n int) []string {
	ret := &[]string{}
	parenthese_combine("", n, 0, ret)
	return *ret
}

func parenthese_combine(s string, l, r int, ret *[]string) {
	if l == 0 && r == 0 {
		*ret = append(*ret, s)
		return
	}

	if l > 0 {
		parenthese_combine(s+"(", l-1, r+1, ret)
	}

	if r > 0 {
		parenthese_combine(s+")", l, r-1, ret)
	}

}
