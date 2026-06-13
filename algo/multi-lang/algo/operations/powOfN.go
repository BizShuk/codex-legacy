package op

// bit manipulate
func PowOfTwo(n int) bool {
	return (n > 0 && (n&(n-1)) == 0)
}

//    return ( n>0 &&  1162261467%n==0);
//	32bits中 3^20 最大的3的n次方
func PowOfThree(n int) bool {
	return (n > 0 && (n == 1 || (n%3 == 0 && PowOfThree(n/3))))
}

// Proof #1: (4^n-1) = (4-1) (4^(n-1) + 4^(n-2) + 4^(n-3) + ..... + 4 + 1)
// Proof #2 (by induction) 4^(n+1) - 1 = 44^n -1 = 34^n + 4^n-1. The first is divided by 3, the second is proven by induction hypothesis
func PowOfFour_math(n int) bool {
	return n > 0 && (n&(n-1)) == 0 && (n-1)%3 == 0
}

func PowOfFour_iter(n int) bool {
	if n == 0 {
		return false
	}
	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}
