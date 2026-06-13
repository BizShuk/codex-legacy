package fundation

import (
	"bytes"
	"fmt"
	"unicode"
)

// [Pattern]: [Go char] byte(uint8) from s[i] vs rune(int32) from foreach
// [Special Rune] https://go.dev/ref/spec#Rune_literals
// \a   U+0007 alert or bell
// \b   U+0008 backspace
// \f   U+000C form feed
// \n   U+000A line feed or newline
// \r   U+000D carriage return
// \t   U+0009 horizontal tab
// \v   U+000B vertical tab
// \\   U+005C backslash
// \'   U+0027 single quote  (valid escape only within rune literals)
// \"   U+0022 double quote  (valid escape only within string literals)
func CharMain() {
	s := "1234abcdABCD"
	// s[] is byte == uint
	// for i,e := range s,   e is rune = int32

	bytes_var := []byte(s)

	bytes_var = bytes.ToLower(bytes_var)
	rune_char := rune(bytes_var[0])
	_ = byte(rune_char)

	print(unicode.IsLower(rune_char) || unicode.IsNumber(rune_char))
}

func CharSlice() {
	s := "1234567890abcdefghijklmnopqrstuvwxyz"
	byteSlice := []byte(s)
	runeSlice := []rune(s)

	for i := 0; i < 10; i++ {
		byteSlice[i] += 1
		runeSlice[i] += 1
	}

	fmt.Println(string(byteSlice), byteSlice)
	fmt.Println(string(byteSlice), runeSlice)
}
