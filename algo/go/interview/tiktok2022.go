package i

import "github.com/bizshuk/algo/tree"

// [Question]: A tree can be reset to magic tree.
// What is magic tree?
// It's a binary tree follow rules below
// 1. root value = 0
// 2. left  child = (2 * parent value) +1
// 3. right child = (2 * parent value) +2

// And, create a function to find whether given number is existing in the magic tree

func TikTok_Singapore_2022_round1() {
	root := &tree.Node{}
	ResetMagicTree(root)
	FindTarget(8, root)
}

func ResetMagicTree(node *tree.Node) {

}

func FindTarget(val int, root *tree.Node) bool {
	return false
}

// [Question]: find minimun substring in a with all characters in b
// [Time Complexity]: O(m + n)

// func main() {
// 	TikTok_Singapore_2022_round2("tiktok", "tt")
// 	TikTok_Singapore_2022_round2("tiktok_is_hiring", "lol")
// 	TikTok_Singapore_2022_round2("tiktok", "tiktok")
// }

func TikTok_Singapore_round2(a, b string) string {
	counterB := map[byte]int{}

	for i := 0; i < len(b); i++ {
		counterB[b[i]] += 1
	}

	lp, rp := 0, 0
	counterA := map[byte]int{}
	minSubstring := ""

	for rp < len(a) {
		counterA[a[rp]] += 1
		rp += 1

		if checkCounter(counterA, counterB) {

			// Two conditions:
			// 1. left pointer char is not related to B string
			// 2. left pointer char is related to B string, but more than B string counter+1
			for counterB[a[lp]] == 0 || counterB[a[lp]] > 0 && counterA[a[lp]]-1 >= counterB[a[lp]] {
				counterA[a[lp]] -= 1
				lp += 1
			}

			if len(minSubstring) == 0 || len(minSubstring) > rp-lp {
				minSubstring = a[lp:rp]
			}
		}
	}

	return minSubstring
}

func checkCounter(counterA map[byte]int, counterB map[byte]int) bool {
	for k, v := range counterB {
		if counterA[k] < v {
			return false
		}
	}
	return true
}

// [Question]: System Design, Video Comment Service
func TikTok_Singapore_2022_round3() {

}
