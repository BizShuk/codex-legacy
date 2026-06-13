package dp

import "github.com/bizshuk/algo/util"

// LCS, Longest Common Subsequence

// [Pattern]: [LCS DP] find longest common subsequence in two strings.
// ex: bcde , ace
// [Logic]:
// b != a => Max(sub(cde,ace), sub(bcde,ce))
// c == c => sub(de, e)
func LCS(text1 string, text2 string) int {
	// [Hint]: extra space for empty string with text2
	dp := make([][]int, len(text1)+1)

	for i := range dp { // [Hint]: extra space for empty string with text1
		dp[i] = make([]int, len(text2)+1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = util.Max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[0][0]
}
