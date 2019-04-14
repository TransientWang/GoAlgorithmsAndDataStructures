package DynamicPlanning

import "math"

func LongestPalindromeSubseq(s string) int {
	size := len(s)

	if size == 0 || size == 1 {
		return size
	}
	var dp = make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
	}
	for i := size - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < size; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = int(math.Max(float64(dp[i+1][j]), float64(dp[i][j-1])))
			}
		}
	}
	return dp[0][size-1]

}
