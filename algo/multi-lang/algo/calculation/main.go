package main

import (
	"fmt"
)

// 1+2+3+...+n = ?
func main() {
	fmt.Printf("sumNfrom1 forloop:%d\n", sumNfrom1_1(30))
	fmt.Printf("sumNfrom1 recursive:%d\n", sumNfrom1_2(30))
	fmt.Printf("sumNfrom1 math:%d\n", sumNfrom1_3(30))
}

// for loop
func sumNfrom1_1(n int) int {
	result := 0
	for i := 1; i < n+1; i++ {
		result += i
	}
	return result
}

// recursive
func sumNfrom1_2(n int) int {
	if n == 1 {
		return 1
	}
	return n + sumNfrom1_2(n-1)
}

// math
func sumNfrom1_3(n int) int {
	return (1 + n) * n / 2
}
