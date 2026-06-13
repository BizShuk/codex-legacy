package base_convert

import (
	"fmt"
	"math"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func convertToTitle(n int) string {
	var b []byte
	for ; n > 0; n = (n - 1) / 26 {
		b = append(b, letters[(n-1)%26])
	}
	fmt.Println(string(b))
	return string(b)
}

func convertToTitle2(n int) string {
	var result string = ""

	for ; n > 0; n = (n - 1) / 26 {
		result = string(letters[(n-1)%26]) + result
	}
	fmt.Println(result)
	return result
}

func titleToNumber(s string) int {
	charmap := map[string]float64{
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5,
		"F": 6, "G": 7, "H": 8, "I": 9, "J": 10,
		"K": 11, "L": 12, "M": 13, "N": 14, "O": 15,
		"P": 16, "Q": 17, "R": 18, "S": 19, "T": 20,
		"U": 21, "V": 22, "W": 23, "X": 24, "Y": 25,
		"Z": 26,
	}

	var y float64 = 0
	var x float64 = 26
	var result float64 = 0
	for i := len(s) - 1; i >= 0; i-- {
		result += charmap[string(s[i])] * math.Pow(x, y)
		y++
	}
	fmt.Println(int(result))
	return int(result)
}
