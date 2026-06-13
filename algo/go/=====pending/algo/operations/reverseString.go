package op

func ReverseString(s string) string {
	slen := len(s)
	var ret = make([]byte, slen)
	for i := 0; i < slen; i++ {
		ret[i] = s[slen-i-1]
	}
	return string(ret)
}

// BUG() ,No extra space
func ReverseString_nospace(s string) string {
	slen := len(s)
	sb := []byte(s)
	for i := 0; i < slen/2; i++ {
		sb[i], sb[slen-1-i] = sb[slen-1-i], sb[i]
	}
	return string(sb)
}
