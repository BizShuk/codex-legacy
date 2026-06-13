package fundation

import (
	"fmt"
	"regexp"
	"strings"
)

// [tip]: strings.Split
func StringSplit_1(s string) {
	fmt.Printf("input string:\n%v", s)

	fmt.Println("\nresult:")
	splited := strings.Split(s, "\n")
	for i, line := range splited {
		fmt.Printf("line %v: %v\n", i, line)
	}
}

// [tip]: strings.Fields remove any space char including new line
func StringSplit_2(s string) {
	fmt.Printf("input string:\n%v", s)

	fmt.Println("\nresult:")
	splited := strings.Fields(s)
	for i, token := range splited {
		fmt.Printf("token %v: %v\n", i, token)
	}
}

// [tip]: It splits a string into substrings separated by a regular expression. The method takes an integer argument n; if n >= 0, it returns at most n substrings.
func StringSplit_3(s string) {
	fmt.Printf("input string:\n%v", s)

	fmt.Println("\nresult:")
	a := regexp.MustCompile(`\n`)
	fmt.Printf("%q\n", a.Split(s, -1))
	fmt.Printf("%q\n", a.Split(s, 0))
	fmt.Printf("%q\n", a.Split(s, 1))
	fmt.Printf("%q\n", a.Split(s, 2))
}
