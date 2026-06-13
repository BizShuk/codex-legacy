package bit

// Find single number, all are duplicate once except one

// [Pattern]: [XOR(^)] true-false
// XOR is
// true , true  => false
// true , false => true
// false, true  => true
// false, false => false

// If same bit => reset to 0
// [Pattern]: [xor op] find unique number
func Xor(nums []int) int {
	xor_result := 0

	for _, num := range nums {
		xor_result ^= num
	}
	return xor_result
}
